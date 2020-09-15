package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	// 검색을 원하는 그룹아이디
	groupId := "G4V20200819094043TGFVDDSDZXJ7O7H"

	// 검색조건값
	params := make(map[string]string)
	params["limit"] = "1"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.GetGroupMessageList(groupId, params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
