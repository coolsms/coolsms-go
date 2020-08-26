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

	// 메시지 아이디로 조회
	// params["criteria"] = "messageId"
	// params["cond"] = "eq"
	// params["value"] = "M4V20200826154526YBEPH9DUHEWWTAZ"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.GetMessageList(params)
	if err != nil {
		fmt.Println(err)
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
