/* validator.go */

package main

import (
	"strings"
)

type Message struct {
	Item   string
	Errors map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Item) == "" {
		msg.Errors["Item"] = "Please enter an item"
	}

	return len(msg.Errors) == 0
}

