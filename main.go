package main

import (
	_ "embed"
	"log/slog"
	"net/http"
)

var endOfDay = 21
var startOfShift = 8
var endOfShift = 16

func main() {
	server := initServer()

	go runExpenseSchedule()
	go runMessageSchedule()

	slog.Info("Serving on http://localhost:12847")
	http.ListenAndServe("localhost:12847", server)
}
