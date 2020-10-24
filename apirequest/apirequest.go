package apirequest

import (
	"bytes"
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/coolsms/coolsms-go/types"
)

const sdkVersion string = "GO-SDK v1.0"

var (
	errFailedToConvertJSON   = errors.New("FailedToConvertJSON")
	errFailedToClientRequest = errors.New("FailedToClientRequest")
)

// APIRequest api
type APIRequest struct {
	// HTTP Request response, statusCode
	response   string `json:"response"`
	statusCode string `json:"statusCode"`

	// Config
	APIKey     string `json:"apiKey"`
	APISecret  string `json:"APISecret"`
	Protocol   string `json:"Protocol"`
	Domain     string `json:"Domain"`
	Prefix     string `json:"Prefix"`
	AppId      string `json:"AppId"`
	SdkVersion string `json:"sdkVersion"`
	OsPlatform string `json:"osPlatform"`

	// Custom Config
	Config map[string]string
}

// RandomString returns a random string
func RandomString(n int) string {
	b := make([]byte, n)
	_, _ = cr.Read(b)

	return hex.EncodeToString(b)
}

// Exists file check
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Get config file path
func getConfigFilePath() string {
	pathList := [5]string{
		"config.json",
		"../config.json",
		"../../config.json",
		"../../../config.json",
		"../../../../config.json",
	}

	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Dir(b)
	filePath = filepath.Join(filePath, "../config.json")

	path, err := os.Getwd()
	if err == nil {
		for _, configPath := range pathList {
			processFilePath := filepath.Join(path, configPath)
			exist, _ := exists(processFilePath)
			if exist == true {
				filePath = processFilePath
				break
			}
		}
	}

	return filePath
}

// NewAPIRequest create
func NewAPIRequest() *APIRequest {
	goos := runtime.GOOS
	goVersion := runtime.Version()
	osPlatform := fmt.Sprintf("%s/%s", goos, goVersion)

	request := APIRequest{response: "", statusCode: "", OsPlatform: osPlatform, SdkVersion: sdkVersion}

	// Read File
	filePath := getConfigFilePath()
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln("Error reading")
		return &request
	}

	err = json.Unmarshal(file, &request)
	fmt.Println(request)
	if err != nil {
		log.Fatalln("Error file Unmarshal")
		return &request
	}

	return &request
}

// SetCustomConfig set custom config for request
func (a *APIRequest) SetCustomConfig(config map[string]string) error {
	for key, value := range config {
		switch key {
		case "APIKey":
			a.APIKey = value
			break
		case "APISecret":
			a.APISecret = value
			break
		case "Protocol":
			a.Protocol = value
			break
		case "Domain":
			a.Domain = value
			break
		case "Prefix":
			a.Prefix = value
			break
		case "AppId":
			a.AppId = value
		}
	}
	return nil
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
	url := fmt.Sprintf("%s://%s/%s%s", a.Protocol, a.Domain, a.Prefix, resource)
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
		return errFailedToClientRequest
	}

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		err = json.NewDecoder(resp.Body).Decode(&errorStruct)
		if err != nil {
			return err
		}
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	err = json.NewDecoder(resp.Body).Decode(&customStruct)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Request method request
func (a *APIRequest) Request(method string, resource string, params interface{}, customStruct interface{}) error {
	// Convert to json string
	jsonString, err := json.Marshal(params)
	if err != nil {
		return errFailedToConvertJSON
	}

	// Prepare for Http Request
	client := &http.Client{}
	url := fmt.Sprintf("%s://%s/%s%s", a.Protocol, a.Domain, a.Prefix, resource)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonString))

	// Set Headers
	authorization := a.GetAuthorization()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	// Request
	resp, err := client.Do(req)
	if err != nil {
		return errFailedToClientRequest
	}

	// StatusCode가 200이 아니라면 에러로 처리
	if resp.StatusCode != 200 {
		errorStruct := types.CustomError{}
		err = json.NewDecoder(resp.Body).Decode(&errorStruct)
		if err != nil {
			return err
		}
		errString := fmt.Sprintf("%s[%d]:%s", errorStruct.ErrorCode, resp.StatusCode, errorStruct.ErrorMessage)
		return errors.New(errString)
	}

	err = json.NewDecoder(resp.Body).Decode(&customStruct)
	if err != nil {
		return err
	}
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
