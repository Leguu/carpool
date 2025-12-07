package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type State struct {
	Going     bool
	Returning bool
}

func (s State) Save() {
	res, _ := json.Marshal(s)
	os.WriteFile("state.json", res, 0644)
}

func (s *State) Load() {
	res, err := os.ReadFile("state.json")
	if err == nil {
		json.Unmarshal(res, &s)
	}
}

var currentState = State{}

func init() {
	currentState.Load()
}

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

	return fmt.Sprintf("I will be %s tomorrow", strings.Join(messages, " and "))
}
