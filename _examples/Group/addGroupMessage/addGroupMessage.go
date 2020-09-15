package main

import (
	"fmt"

	"github.com/coolsms/coolsms-go"
)

func main() {
	client := coolsms.NewClient()

	// 메시지를 추가할 그룹 아이디
	groupId := "G4V20200824233846LAI57B7SKJACJ40"

	// 추가할 메시지 데이터
	message := make(map[string]interface{})
	message["to"] = "01000000000"
	message["from"] = "029302266"
	message["text"] = "안녕하세요 홍길동님 회원가입을 환영합니다."
	message["type"] = "SMS"

	// 포토문자(MMS)
	// Storage 예제에서 이미지 파일 업로드 방법을 확인하실 수 있습니다.
	// message["type"] = "MMS"
	// message["imageId"] = "ST01FJ20073019363R8G58YLK1kvUXEH"
	// message["subject"] = "Subject Title"

	// 장문문자(LMS)
	// message["type"] = "LMS"
	// message["subject"] = "Subject Title"

	// 친구톡(CTA)
	// 사이트에서 플러스친구 연동 후 사용이 가능합니다.
	// message["type"] = "CTA"
	// kakaoOptions := make(map[string]interface{})
	// kakaoOptions["pfId"] = "KA01PF2003231823W4979UIEbV3fHtkY"
	// message["kakaoOptions"] = kakaoOptions

	// 알림톡(CTA)
	// 사이트에서 플러스친구 연동 및 템플릿 등록 후 사용이 가능합니다.
	// message["type"] = "ATA"
	// kakaoOptions := make(map[string]interface{})
	// kakaoOptions["pfId"] = "KA01PF2003231823W4979UIEbV3fHtkY"
	// kakaoOptions["templateId"] = "KA01TP2003W3882345595hMWrRtFCSTq"
	// message["kakaoOptions"] = kakaoOptions

	// 메시지 데이터를 목록화 하여 파라미터로 넘겨줍니다.
	var messageList []interface{}
	messageList = append(messageList, message)

	fmt.Println(messageList)

	params := make(map[string]interface{})
	params["messages"] = messageList

	// API 호출 후 결과값을 받아 옵니다.
	result, err := client.Messages.AddGroupMessage(groupId, params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print Result
	fmt.Printf("%+v\n", result)
}
