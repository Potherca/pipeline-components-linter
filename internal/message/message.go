package message

import (
	"fmt"
	"internal/check"
)

type Marker struct {
	Fail       string
	Incomplete string
	Pass       string
	Skip       string
}

func CreateMessage(status check.Status, code string, message string, messageMarker Marker) string {
	var marker string

	switch status {
	case check.Pass:
		marker = messageMarker.Pass
	case check.Fail:
		marker = messageMarker.Fail
	default:
		panic("unhandled default case")
	}

	return fmt.Sprintf("%s %s %s\n", code, marker, message)
}
