package restclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"os"

	log "github.com/jcelliott/lumber"
)

type Restclient interface {
	OpenGETRequest(url string, record interface{})
	OpenDELETERequest(url string)
	OpenRequestWithInput(url string, method string, input interface{}, response interface{})
}

type restclientImpl struct {
	baseUrl                 string
	basicAuthorizationValue string
}

func CreateRestclient(baseUrl, username, password string) Restclient {
	authToken := []byte(fmt.Sprintf("%s:%s", username, password))
	authTokenHash := base64.StdEncoding.EncodeToString(authToken)
	return &restclientImpl{baseUrl: baseUrl, basicAuthorizationValue: authTokenHash}
}

func (endpoint *restclientImpl) OpenGETRequest(url string, record interface{}) {
	endpoint.openRequestWithEncodedInput(url, "GET", new(bytes.Buffer), record)
}

func (endpoint *restclientImpl) OpenDELETERequest(url string) {
	endpoint.openRequestWithEncodedInput(url, "DELETE", new(bytes.Buffer), nil)
}

func (endpoint *restclientImpl) OpenRequestWithInput(url string, method string, input interface{}, record interface{}) {
	// Encode Body
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&input)

	endpoint.openRequestWithEncodedInput(url, method, body, record)
}

func (endpoint *restclientImpl) openRequestWithEncodedInput(url string, method string, bufferedInput *bytes.Buffer, record interface{}) {
	log.Debug("Open %s %s Request for \"%s\"", method, endpoint.baseUrl, url)

	// Build the request
	req, err := http.NewRequest(method, endpoint.baseUrl+url, bufferedInput)
	if err != nil {
		log.Fatal("Error init Request: ", err)
		os.Exit(-1)
	}

	req.Header.Set("X-Atlassian-Token", "no-check")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", endpoint.basicAuthorizationValue))

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending Request: ", err)
		os.Exit(-1)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	if !strings.HasPrefix(resp.Status, "2") {
		log.Fatal("Request not successful: Http %d", resp.StatusCode)
		os.Exit(-1)
	}

	if resp.StatusCode == 204 {
		return
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading Response: ", err)
		os.Exit(-1)
	}

	payload := string(responseData)

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(strings.NewReader(payload)).Decode(record); err != nil {
		log.Warn("DecodingError: ", err)
	}
}
