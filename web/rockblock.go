package web

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/SmartParksOrg/smartparks-connect-web/utils"
)

type RockBLOCKRequest struct {
	RequestType string `json:"request_type"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        uint8  `json:"port"`
	IMEI        string `json:"imei"`
	Confirmed   bool   `json:"confirmed"`

	ID      string      `json:"id"`
	Length  uint8       `json:"content_length"`
	Type    string      `json:"content_type"`
	Content interface{} `json:"content"`
}

type RockBLOCKLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) handleRockBLOCKLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var request RockBLOCKLoginRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new HTTP POST request with the RockBLOCK API endpoint URL
	url := fmt.Sprintf("https://rockblock.rock7.com/rockblock/MT?username=%s&password=%s&data=0",
		url.QueryEscape(request.Username),
		url.QueryEscape(request.Password),
	)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("accept", "text/plain")

	// Send the request and parse the response
	res, err := h.httpClient.Do(req)
	if err != nil {
		Resp(w, nil, err)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// Check the response for errors
	if strings.HasPrefix(string(body), "FAILED") {
		// Response is an error, get the code from the response message
		parts := strings.SplitN(string(body), ",", 3)
		code, _ := strconv.Atoi(parts[1])
		if code == 10 {
			Resp(w, nil, errors.New(parts[2]))
			return
		}
		// Other error codes are ignored at this stage, since we are not actually sending data to any device, we just
		// want to check if credentials are correct
	}

	Resp(w, nil, nil)
}

func (h *Handler) handleRockBLOCKQueue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON data
	decoder := json.NewDecoder(r.Body)
	var request RockBLOCKRequest
	err := decoder.Decode(&request)
	if err != nil {
		Resp(w, nil, errors.New("Invalid request body: "+err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if request.Port == 0 || request.ID == "" || request.Type == "" || request.Length == 0 {
		// w.WriteHeader(http.StatusBadRequest)
		Resp(w, nil, errors.New("missing one or more of required parameters: port, content_type, id, type"))
		return
	}

	byteData, err := utils.ConvertRockBLOCKBytes(request.Port, request.ID, utils.VType(request.Type), request.Length, request.Content)
	if err != nil {
		// w.WriteHeader(500)
		Resp(w, nil, err)
		return
	}

	// We reuse confirmed from Chirpstack as Flush value
	flush := "no"
	if request.Confirmed {
		flush = "yes"
	}

	// Create a new HTTP POST request with the RockBLOCK API endpoint URL
	url := fmt.Sprintf("https://rockblock.rock7.com/rockblock/MT?username=%s&password=%s&data=%s&flush=%s&imei=%s",
		url.QueryEscape(request.Username),
		url.QueryEscape(request.Password),
		url.QueryEscape(hex.EncodeToString(byteData)),
		url.QueryEscape(flush),
		url.QueryEscape(request.IMEI),
	)
	client := h.httpClient
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("accept", "text/plain")

	// Send the request and parse the response
	res, err := client.Do(req)
	if err != nil {
		Resp(w, nil, err)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// Check the response for errors
	if strings.HasPrefix(string(body), "FAILED") {
		// Response is an error, get the code from the response message
		parts := strings.SplitN(string(body), ",", 3)

		Resp(w, map[string]interface{}{"response": string(body), "request_data": hex.EncodeToString(byteData)}, errors.New(parts[2]))
		return
	}

	parts := strings.SplitN(string(body), ",", 2)
	Resp(w, map[string]interface{}{"mtId": parts[1], "Payload": hex.EncodeToString(byteData), "Base64": base64.RawStdEncoding.EncodeToString(byteData)}, err)
}
