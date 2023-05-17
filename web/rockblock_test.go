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
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
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
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
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
