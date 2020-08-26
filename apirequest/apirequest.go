package apirequest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/coolsms/coolsms-go/types"
)

const sdkVersion string = "GO-SDK v1.0"

var errFailedToConvertJSON = errors.New("FailedToConvertJSON")
var errFailedToClientRequest = errors.New("FailedToClientRequest")
var errFailedToReadFile = errors.New("FailedToReadFile")

// APIRequest api
type APIRequest struct {
	// HTTP Request response, statusCode
	response   string
	statusCode string

	// Config
	APIKey     string `json:"apiKey"`
	APISecret  string `json:"APISecret"`
	Protocol   string `json:"Protocol"`
	Domain     string `json:"Domain"`
	Prefix     string `json:"Prefix"`
	AppId      string `json: "AppId"`
	SdkVersion string
	OsPlatform string
}

// RandomString returns a random string
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// NewAPIRequest create
func NewAPIRequest() *APIRequest {
	os := runtime.GOOS
	goVersion := runtime.Version()
	osPlatform := fmt.Sprintf("%s/%s", os, goVersion)

	request := APIRequest{response: "", statusCode: "", OsPlatform: osPlatform, SdkVersion: sdkVersion}

	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Dir(b)
	file, err := ioutil.ReadFile(filepath.Join(filePath, "../config.json"))
	if err != nil {
		log.Fatalln("Error reading")
		return &request
	}

	json.Unmarshal([]byte(file), &request)
	return &request
}

// GetAuthorization gets the authorization
func (a *APIRequest) GetAuthorization() string {
	salt := RandomString(20)
	date := time.Now().Format(time.RFC3339)
	h := hmac.New(sha256.New, []byte(a.APISecret))
	h.Write([]byte(date + salt))
	signature := hex.EncodeToString(h.Sum(nil))
	authorization := fmt.Sprintf("HMAC-SHA256 apiKey=%s, date=%s, salt=%s, signature=%s", a.APIKey, date, salt, signature)
	return authorization
}

// GET method request
func (a *APIRequest) GET(resource string, params map[string]string, customStruct interface{}) error {
	// Prepare for Http Request
	client := &http.Client{}
	url := fmt.Sprintf("%s://%s/%s", a.Protocol, a.Domain, resource)
	fmt.Println("URL :", url)
	req, _ := http.NewRequest("GET", url, nil)

	// Set Query Parameters
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Set Headers
	authorization := a.GetAuthorization()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errFailedToClientRequest
	}

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		json.NewDecoder(resp.Body).Decode(&errorStruct)
		fmt.Println("StatusCodeError:", resp.StatusCode)
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	json.NewDecoder(resp.Body).Decode(&customStruct)
	defer resp.Body.Close()
	return nil
}

// Request method request
func (a *APIRequest) Request(method string, resource string, params interface{}, customStruct interface{}) error {
	// Convert to json string
	jsonString, jsonErr := json.Marshal(params)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return errFailedToConvertJSON
	}

	// Prepare for Http Request
	client := &http.Client{}
	url := fmt.Sprintf("%s://%s/%s", a.Protocol, a.Domain, resource)
	fmt.Println("URL :", url)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonString))

	// Set Headers
	authorization := a.GetAuthorization()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return errFailedToClientRequest
	}

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		json.NewDecoder(resp.Body).Decode(&errorStruct)
		fmt.Println("StatusCodeError:", resp.StatusCode)
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	json.NewDecoder(resp.Body).Decode(&customStruct)
	defer resp.Body.Close()

	return nil
}

// POST method request
func (a *APIRequest) POST(resource string, params interface{}, customStruct interface{}) error {
	err := a.Request("POST", resource, params, &customStruct)
	return err
}

// PUT method request
func (a *APIRequest) PUT(resource string, params interface{}, customStruct interface{}) error {
	err := a.Request("PUT", resource, params, &customStruct)
	return err
}

// DELETE method request
func (a *APIRequest) DELETE(resource string, params interface{}, customStruct interface{}) error {
	err := a.Request("DELETE", resource, params, &customStruct)
	return err
}
