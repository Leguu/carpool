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
		if time.Now().Hour() < endOfShift {
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
		currentState.Save()

		if endOfDay <= hour() || hour() < startOfShift {
			sendMessageToLegu(fmt.Sprintf("I have changed my mind, I will %sbe going with you %s.", not(currentState.Going), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	mux.HandleFunc("/api/state/returning", func(w http.ResponseWriter, r *http.Request) {
		currentState.Returning = !currentState.Returning
		currentState.Save()

		if endOfDay <= hour() || hour() < startOfShift {
			sendMessageToLegu(fmt.Sprintf("I have changed my mind, I will %sbe returning with you %s.", not(currentState.Returning), day()))
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	})

	index := func(w http.ResponseWriter, r *http.Request) {
		page := views.IndexPage{
			Day:        day(),
			Going:      currentState.Going,
			Returning:  currentState.Returning,
			EndOfShift: endOfShift,
			EndOfDay:   endOfDay,
			StartOfDay: startOfShift,
		}
		if startOfShift <= hour() && hour() < endOfShift {
			page.DisableGoing = true
		} else {
			page.DisableGoing = false
		}
		result := views.RenderIndex(page)
		w.Write([]byte(result))
	}
	mux.HandleFunc("/", index)
	mux.HandleFunc("/*", index)

	return mux
}
