package main

import (
	_ "embed"
	"log/slog"
	"net/http"
)

func main() {
	server := initServer()

	go runSchedule()

	slog.Info("Serving on http://localhost:12847")
	http.ListenAndServe("localhost:12847", server)
}
