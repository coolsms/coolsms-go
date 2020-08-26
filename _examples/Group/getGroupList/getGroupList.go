package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	// 검색조건값
	params := make(map[string]string)
	params["limit"] = "1"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.GetGroupList(params)
	if err != nil {
		fmt.Println(err)
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
