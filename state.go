package main

import (
	"fmt"
	"strings"
)

type State struct {
	Going        bool
	Returning    bool
	DisableGoing bool
	Day          string
}

var currentState = State{}

func (currentState State) getAmount() int {
	amount := 0
	if currentState.Going {
		amount += 400
	}
	if currentState.Returning {
		amount += 400
	}
	return amount
}

func (currentState State) getMessage() string {
	var messages []string
	if currentState.Going {
		messages = append(messages, "going with you")
	}
	if currentState.Returning {
		messages = append(messages, "returning with you")
	}

	return fmt.Sprintf("Gary will be %s tomorrow", strings.Join(messages, " and "))
}
