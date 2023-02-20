package web

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/SmartParksOrg/smartparks-connect-web/device_template"
	"github.com/SmartParksOrg/smartparks-connect-web/grpc_client"
	"github.com/SmartParksOrg/smartparks-connect-web/utils"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Request struct {
	RequestType string `json:"request_type"`
	Port        int    `json:"port"`
	ServerUrl   string `json:"server_url"`
	ApiKey      string `json:"api_key"`
	DeviceID    string `json:"device_id"`
	Confirmed   bool   `json:"confirmed"`

	ID      string      `json:"id"`
	Length  uint8       `json:"content_length"`
	Type    string      `json:"content_type"`
	Content interface{} `json:"content"`
}
type ListRequest struct {
	ServerUrl string `json:"server_url"`
	ApiKey    string `json:"api_key"`
	ID        int64  `json:"id"`
	ListType  string `json:"list_type"`
}

type LoginRequest struct {
	ServerUrl string `json:"server_url"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var request LoginRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}
	jwt, err := grpc_client.DefaultGrpcClient.Login(r.Context(), request.ServerUrl, request.Email, request.Password)
	Resp(w, jwt, err)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var request ListRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	serverParam := grpc_client.ServerParam{
		ServerUrl: request.ServerUrl,
		ApiToken:  request.ApiKey,
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	var list interface{}
	switch request.ListType {
	case "org":
		list, err = grpc_client.DefaultGrpcClient.GetOrgList(ctx, serverParam)
	case "app":
		list, err = grpc_client.DefaultGrpcClient.GetApplicationList(ctx, serverParam, request.ID)
	case "device":
		list, err = grpc_client.DefaultGrpcClient.GetDeviceList(ctx, serverParam, request.ID)
	}
	Resp(w, list, err)
}
func handleAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var request Request
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}
	byteData, err := utils.ConvertBytes(request.ID, utils.VType(request.Type), request.Length, request.Content)
	if err != nil {
		Err(w, err)
		return
	}
	log.Printf("convert bytes : %s ", hex.EncodeToString(byteData))
	serverParam := grpc_client.ServerParam{
		ServerUrl: request.ServerUrl,
		ApiToken:  request.ApiKey,
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()
	FPort := request.Port
	if FPort == 0 && request.RequestType == "settings" {
		FPort = 3
	} else if FPort == 0 && request.RequestType == "commands" {
		FPort = 32
	}
	FCnt, err := grpc_client.DefaultGrpcClient.DeviceEnqueue(ctx, serverParam, request.DeviceID, uint32(FPort), request.Confirmed, byteData)
	Resp(w, map[string]interface{}{"FCnt": FCnt, "Paylod": hex.EncodeToString(byteData), "Base64": base64.RawStdEncoding.EncodeToString(byteData)}, err)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("public/dist/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func handleTemplateList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Succ(w, device_template.Templates)
}
