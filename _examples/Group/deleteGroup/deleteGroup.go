package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	// 삭제할 그룹 아이디
	groupId := "G4V20200825011330MWECXBWMBIABRKL"

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.DeleteGroup(groupId)
	if err != nil {
		fmt.Println(err)
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
