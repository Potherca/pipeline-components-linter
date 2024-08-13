package message

import (
	"internal/check"
)

type Marker struct {
	Fail       string
	Incomplete string
	Pass       string
	Skip       string
}

type Message struct {
	Code    string
	Message string
	Status  check.Status
}

func CreateMessage(status check.Status, code string, message string) Message {
	return Message{
		Code:    code,
		Message: message,
		Status:  status,
	}
}
