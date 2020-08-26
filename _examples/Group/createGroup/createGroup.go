package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	//  파라미터들
	params := make(map[string]string)

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.CreateGroup(params)
	if err != nil {
		fmt.Println(err)
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
