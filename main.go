package main

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func main() {
	server := initServer()

	go runSchedule()

	sendMessageToLegu("Carpool server started")
	defer sendMessageToLegu("Be warned, carpool server is shutting down")

	slog.Info("Serving on http://localhost:12847")
	http.ListenAndServe("localhost:12847", server)
}
