package messages

import (
	"fmt"

	"github.com/coolsms/coolsms-go/apirequest"
	"github.com/coolsms/coolsms-go/types"
)

// Messages struct
type Messages struct{}

// GetMessageList gets the list of messages
func (r *Messages) GetMessageList(params map[string]string) (types.MessageList, error) {
	request := apirequest.NewAPIRequest()
	result := types.MessageList{}
	err := request.GET("messages/v4/list", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SendSimpleMessage sends a simple message
func (r *Messages) SendSimpleMessage(params map[string]interface{}) (types.SimpleMessage, error) {
	request := apirequest.NewAPIRequest()
	result := types.SimpleMessage{}
	err := request.POST("messages/v4/send", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CreateGroup creeate message group
func (r *Messages) CreateGroup(params map[string]string) (types.Group, error) {
	request := apirequest.NewAPIRequest()
	result := types.Group{}
	err := request.POST("messages/v4/groups", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// AddGroupMessage adds a group message
func (r *Messages) AddGroupMessage(groupId string, params interface{}) (types.AddGroupMessageList, error) {
	request := apirequest.NewAPIRequest()
	result := types.AddGroupMessageList{}
	url := fmt.Sprintf("messages/v4/groups/%s/messages", groupId)
	err := request.PUT(url, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SendGroup send a group
func (r *Messages) SendGroup(groupId string) (types.Group, error) {
	request := apirequest.NewAPIRequest()
	result := types.Group{}
	url := fmt.Sprintf("messages/v4/groups/%s/send", groupId)
	params := make(map[string]string)
	err := request.POST(url, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// DeleteGroup delete a group
func (r *Messages) DeleteGroup(groupId string) (types.Group, error) {
	request := apirequest.NewAPIRequest()
	result := types.Group{}
	url := fmt.Sprintf("messages/v4/groups/%s", groupId)
	params := make(map[string]string)
	err := request.DELETE(url, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroupList gets the list of groups
func (r *Messages) GetGroupList(params map[string]string) (types.GroupList, error) {
	request := apirequest.NewAPIRequest()
	result := types.GroupList{}
	err := request.GET("messages/v4/groups", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroup get a group
func (r *Messages) GetGroup(groupId string) (types.Group, error) {
	request := apirequest.NewAPIRequest()
	result := types.Group{}
	params := map[string]string{}
	url := fmt.Sprintf("messages/v4/groups/%s", groupId)
	err := request.GET(url, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetGroupMessageList returns a list of group messages
func (r *Messages) GetGroupMessageList(groupId string, params map[string]string) (types.MessageList, error) {
	request := apirequest.NewAPIRequest()
	result := types.MessageList{}
	url := fmt.Sprintf("messages/v4/groups/%s/messages", groupId)
	err := request.GET(url, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
