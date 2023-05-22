package web

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	VALID_QUEUE_BODY      = []byte(`{"username":"testuser","Password":"testpass","id":"0x01","imei":"123451234512345","confirmed":false,"request_type":"settings","port":3,"content":0,"content_type":"uint32","content_length":4}`)
	VALID_QUEUE_BODY_2    = []byte(`{"username":"testuser","Password":"testpass","id":"0x01","imei":"123451234512345","confirmed":true,"request_type":"settings","port":3,"content":0,"content_type":"uint32","content_length":4}`)
	INVALID_QUEUE_BODY    = []byte(`{"username":1,"Password":"testpass","id":"0x01","imei":"123451234512345","confirmed":false,"request_type":"settings","port":3,"content":0,"content_type":"uint32","content_length":4}`)
	INVALID_QUEUE_BODY_2  = []byte(`{"username":"testuser","Password":"testpass","id":"đčžćč","imei":"123451234512345","confirmed":true,"request_type":"settings","port":3,"content":0,"content_type":"uint32","content_length":4}`)
	INCOMPLETE_QUEUE_BODY = []byte(`{"username":"testuser","Password":"testpass","confirmed":false,"request_type":"settings","port":3,"content_type":"uint32","content_length":4}`)
)

type mockRoundTripper struct {
	message     string
	status      int
	err         error
	lastRequest *http.Request
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.lastRequest = req

	if m.err != nil {
		return nil, m.err
	}

	return &http.Response{
		StatusCode: m.status,
		Body:       io.NopCloser(strings.NewReader(m.message)),
	}, nil
}

func TestHandleRockBLOCKLoginERROR(t *testing.T) {
	// Create a new JSON object for testing
	requestBody := []byte(`{"username":"testuser","password":"testpass"}`)

	// Test incorrect User
	// Create a new HTTP request and response objects for testing
	req, err := http.NewRequest("POST", "/api/rockblock/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		message: "FAILED,10,Invalid login credentials",
		status:  http.StatusOK,
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	h.handleRockBLOCKLogin(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	expectedResponse := "{\"code\":500,\"result\":null,\"err_msg\":\"Invalid login credentials\"}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Check that the HTTP request made to the RockBLOCK API was correct
	expectedURL := "https://rockblock.rock7.com/rockblock/MT?username=testuser&password=testpass&data=0"
	assert.Equal(t, expectedURL, mockTransport.lastRequest.URL.String())
	assert.Equal(t, "POST", mockTransport.lastRequest.Method)
	assert.Equal(t, "text/plain", mockTransport.lastRequest.Header.Get("accept"))

	// Test wrong method
	req, err = http.NewRequest("GET", "/api/rockblock/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	h.handleRockBLOCKLogin(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Result().StatusCode)

	// Test invalid request parameters
	requestBody = []byte(`{"username":1,"password":"testpass"}`)
	req, err = http.NewRequest("POST", "/api/rockblock/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	h.handleRockBLOCKLogin(rr, req)
	// Check that the expected response was returned
	assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	assert.Contains(t, rr.Body.String(), "Invalid request body:")

	// Test network error
	// Create a mock response from the RockBLOCK API server
	mockTransport = &mockRoundTripper{
		err: errors.New("test network error"),
	}
	client = &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h = Handler{httpClient: client}
	rr = httptest.NewRecorder()

	requestBody = []byte(`{"username":"testuser","password":"testpass"}`)
	req, err = http.NewRequest("POST", "/api/rockblock/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	h.handleRockBLOCKLogin(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	expectedResponse = "{\"code\":500,\"result\":null,\"err_msg\":\"Post \\\"https://rockblock.rock7.com/rockblock/MT?username=testuser\\u0026password=testpass\\u0026data=0\\\": test network error\"}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
}

func TestHandleRockBLOCKLoginOK(t *testing.T) {
	// Test Successfull Login
	// Create a new JSON object for testing
	requestBody := []byte(`{"username":"testuser","password":"testpass"}`)

	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		message: "FAILED,11,No RockBLOCK with this IMEI found on your account",
		status:  http.StatusOK,
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	// Create a new HTTP request and response objects for testing
	req, err := http.NewRequest("POST", "/api/rockblock/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h.handleRockBLOCKLogin(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	expectedResponse := `{"code":200,"result":null}` + "\n"
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Check that the HTTP request made to the RockBLOCK API was correct
	expectedURL := "https://rockblock.rock7.com/rockblock/MT?username=testuser&password=testpass&data=0"
	assert.Equal(t, expectedURL, mockTransport.lastRequest.URL.String())
	assert.Equal(t, "POST", mockTransport.lastRequest.Method)
	assert.Equal(t, "text/plain", mockTransport.lastRequest.Header.Get("accept"))
}

func TestHandleRockBLOCKQueueIncompleteRequestData(t *testing.T) {
	// Test incorrect data
	// Create a new HTTP request and response objects for testing
	req, err := http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(INCOMPLETE_QUEUE_BODY))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		message: "FAILED,10,Invalid login credentials",
		status:  http.StatusOK,
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	h.handleRockBLOCKQueue(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	expectedResponse := "{\"code\":500,\"result\":null,\"err_msg\":\"missing one or more of required parameters: port, content_type, id, type\"}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
}

func TestHandleRockBLOCKQueueInvalidRequestData(t *testing.T) {
	// Test invalid request parameters
	req, err := http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(INVALID_QUEUE_BODY_2))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		message: "FAILED,10,Invalid login credentials",
		status:  http.StatusOK,
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	h.handleRockBLOCKQueue(rr, req)

	h.handleRockBLOCKQueue(rr, req)
	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	assert.Contains(t, rr.Body.String(), "Invalid request body:")
}

func TestHandleRockBLOCKQueueNetworkError(t *testing.T) {
	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		err: errors.New("test network error"),
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	rr := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(VALID_QUEUE_BODY))
	if err != nil {
		t.Fatal(err)
	}

	h.handleRockBLOCKQueue(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	expectedResponse := "{\"code\":500,\"result\":null,\"err_msg\":\"Post \\\"https://rockblock.rock7.com/rockblock/MT?username=testuser\\u0026password=testpass\\u0026data=03010400000000\\u0026flush=no\\u0026imei=123451234512345\\\": test network error\"}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
}

func TestHandleRockBLOCKQueueWrongHTTPMethod(t *testing.T) {
	// Create a mock response from the RockBLOCK API server
	mockTransport := &mockRoundTripper{
		err: errors.New("test network error"),
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	rr := httptest.NewRecorder()

	// Test wrong method
	req, err := http.NewRequest("GET", "/api/rockblock/queue", bytes.NewBuffer(VALID_QUEUE_BODY))
	if err != nil {
		t.Fatal(err)
	}

	h.handleRockBLOCKQueue(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Result().StatusCode)
}

func TestHandleRockBLOCKQueueSuccessful(t *testing.T) {
	// Create a new HTTP request and response objects for testing
	req, err := http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(VALID_QUEUE_BODY))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Mock FAILED response from API
	mockTransport := &mockRoundTripper{
		message: "FAILED,10,Invalid login credentials",
		status:  http.StatusOK,
	}
	client := &http.Client{Transport: mockTransport}

	// Call the handler function with the mock request/response objects
	h := Handler{httpClient: client}
	h.handleRockBLOCKQueue(rr, req)

	// Check that the expected response was returned
	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
	expectedResponse := "{\"code\":500,\"result\":{\"request_data\":\"03010400000000\",\"response\":\"FAILED,10,Invalid login credentials\"},\"err_msg\":\"Invalid login credentials\"}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Check that the HTTP request made to the RockBLOCK API was correct
	expectedURL := "https://rockblock.rock7.com/rockblock/MT?username=testuser&password=testpass&data=03010400000000&flush=no&imei=123451234512345"
	assert.Equal(t, expectedURL, mockTransport.lastRequest.URL.String())
	assert.Equal(t, "POST", mockTransport.lastRequest.Method)
	assert.Equal(t, "text/plain", mockTransport.lastRequest.Header.Get("accept"))

	// Mock OK reponse
	req, err = http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(VALID_QUEUE_BODY))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	mockTransport.message = "OK,10"
	h.handleRockBLOCKQueue(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	expectedResponse = "{\"code\":200,\"result\":{\"Base64\":\"AwEEAAAAAA\",\"Payload\":\"03010400000000\",\"mtId\":\"10\"}}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Mock Second OK reponse
	req, err = http.NewRequest("POST", "/api/rockblock/queue", bytes.NewBuffer(VALID_QUEUE_BODY_2))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	mockTransport.message = "OK,10"
	h.handleRockBLOCKQueue(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	expectedResponse = "{\"code\":200,\"result\":{\"Base64\":\"AwEEAAAAAA\",\"Payload\":\"03010400000000\",\"mtId\":\"10\"}}\n"
	assert.Equal(t, expectedResponse, rr.Body.String())
	expectedURL = "https://rockblock.rock7.com/rockblock/MT?username=testuser&password=testpass&data=03010400000000&flush=yes&imei=123451234512345"
	assert.Equal(t, expectedURL, mockTransport.lastRequest.URL.String())

}
