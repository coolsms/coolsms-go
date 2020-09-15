package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	/*
		업로드할 데이터
		타입별 제한 사이즈
		[KAKAO] : 500KB
		[MMS] : 200KB
		[DOCUMENT] : 2MB
		[COOLSMS-MMS] : 300KB
	*/
	params := make(map[string]string)
	params["file"] = "./test.jpg"
	params["name"] = "customFileName"
	params["type"] = "MMS"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Storage.UploadFile(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
