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

func runMessageSchedule() {
	for {
		time.Sleep(time.Until(nextHour(endOfDay, true)))

		if !currentState.Going && !currentState.Returning {
			continue
		}

		sendMessageToLegu(currentState.getMessage())
	}
}

func runExpenseSchedule() {
	for {
		time.Sleep(time.Until(nextHour(endOfShift, true)))

		if !currentState.Going && !currentState.Returning {
			continue
		}

		slog.Info("Going or returning detected, adding expense")

		addExpense(currentState.getAmount())

		currentState = State{}
		currentState.Save()
	}
}
