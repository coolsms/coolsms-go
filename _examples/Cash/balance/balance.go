package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Cash.Balance()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
