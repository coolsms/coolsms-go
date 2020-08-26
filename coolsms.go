package coolsms

import (
	"github.com/coolsms/coolsms-go/cash"
	"github.com/coolsms/coolsms-go/messages"
	"github.com/coolsms/coolsms-go/storage"
)

// Client struct
type Client struct {
	Messages messages.Messages
	Storage  storage.Storage
	Cash     cash.Cash
}

// NewClient return a new client
func NewClient() *Client {
	client := Client{}
	return &client
}
