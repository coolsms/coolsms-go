package storage

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"

	"github.com/coolsms/coolsms-go/apirequest"
	"github.com/coolsms/coolsms-go/types"
)

var (
	errFailToReadFile = errors.New("FailToReadFile")
	errFileNotFound   = errors.New("FileNotFound")
)

// Storage struct
type Storage struct{}

// UploadFile upload a file
func (r *Storage) UploadFile(params map[string]string) (types.File, error) {
	result := types.File{}

	// 파일이 없다면 에러
	if _, ok := params["file"]; !ok {
		return result, errFileNotFound
	}

	// Open file
	f, err := os.Open(params["file"])
	if err != nil {
		return result, errFileNotFound
	}

	// Read entire contents into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return result, errFailToReadFile
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to params.
	params["file"] = encoded

	request := apirequest.NewAPIRequest()
	err = request.POST("storage/v1/files", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetFileList gets the list of files
func (r *Storage) GetFileList(params map[string]string) (types.FileList, error) {
	request := apirequest.NewAPIRequest()
	result := types.FileList{}
	err := request.GET("storage/v1/files", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
