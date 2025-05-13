package main

import (
	"carpool/views"
	"fmt"
	"net/http"
	"time"
)

func initServer() *http.ServeMux {
	mux := http.NewServeMux()

	not := func(n bool) string {
		if !n {
			return "NOT "
		} else {
			return ""
		}
	}
	day := func() string {
		if time.Now().Hour() <= 8 {
			return "today"
		} else {
			return "tomorrow"
		}
	}

	mux.HandleFunc("/api/state/going", func(w http.ResponseWriter, r *http.Request) {
		currentState.Going = !currentState.Going

		if time.Now().Hour() >= 20 || time.Now().Hour() < 8 {
			sendMessageToLegu(fmt.Sprintf("Gary has changed his mind, he will %sbe going with you %s.", not(currentState.Going), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	mux.HandleFunc("/api/state/returning", func(w http.ResponseWriter, r *http.Request) {
		currentState.Returning = !currentState.Returning

		if time.Now().Hour() >= 20 || time.Now().Hour() < 8 {
			sendMessageToLegu(fmt.Sprintf("Gary has changed his mind, he will %sbe going with you %s.", not(currentState.Returning), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	index := func(w http.ResponseWriter, r *http.Request) {
		result := views.IndexTemplate.MustExec(currentState)
		w.Write([]byte(result))
	}
	mux.HandleFunc("/", index)
	mux.HandleFunc("/*", index)

	return mux
}
