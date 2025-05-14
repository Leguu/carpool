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
		if time.Now().Hour() <= endOfShift {
			return "today"
		} else {
			return "tomorrow"
		}
	}

	hour := func() int {
		return time.Now().Hour()
	}

	mux.HandleFunc("/api/state/going", func(w http.ResponseWriter, r *http.Request) {
		if startOfShift <= hour() && hour() < endOfShift {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You can't change the going state right now"))
			return
		}

		currentState.Going = !currentState.Going

		if endOfDay <= hour() || hour() < startOfShift {
			sendMessageToLegu(fmt.Sprintf("Gary has changed his mind, he will %sbe going with you %s.", not(currentState.Going), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	mux.HandleFunc("/api/state/returning", func(w http.ResponseWriter, r *http.Request) {
		currentState.Returning = !currentState.Returning

		if endOfDay <= hour() || hour() < startOfShift {
			sendMessageToLegu(fmt.Sprintf("Gary has changed his mind, he will %sbe returning with you %s.", not(currentState.Returning), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	index := func(w http.ResponseWriter, r *http.Request) {
		ctx := map[string]any{
			"day":        day(),
			"going":      currentState.Going,
			"returning":  currentState.Returning,
			"endOfShift": endOfShift,
			"endOfDay":   endOfDay,
			"startOfDay": startOfShift,
		}
		if startOfShift <= hour() && hour() < endOfShift {
			ctx["disableGoing"] = true
		} else {
			ctx["disableGoing"] = false
		}
		result := views.IndexTemplate.MustExec(ctx)
		w.Write([]byte(result))
	}
	mux.HandleFunc("/", index)
	mux.HandleFunc("/*", index)

	return mux
}
