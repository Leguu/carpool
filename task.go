package main

import (
	"log/slog"
	"time"
)

func nextHour(hour int, skipToday bool) time.Time {
	now := time.Now()
	result := time.Date(
		now.Year(), now.Month(), now.Day(),
		hour, 0, 0, 0, now.Location(),
	)
	if now.After(result) && skipToday {
		result = result.Add(24 * time.Hour)
	}
	return result
}

func runSchedule() {
	for {
		nextNine := nextHour(8, true)
		time.Sleep(time.Until(nextNine))

		if !currentState.Going && !currentState.Returning {
			return
		}

		slog.Info("Going or returning detected, adding expense")

		addExpense(currentState.getAmount())

		sendMessageToLegu(discord, currentState.getMessage())

		currentState = State{}
	}
}
