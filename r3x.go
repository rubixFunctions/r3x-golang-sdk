package r3x

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string
}

func Execute () {
	HTTPStream()
}

func HTTPStream(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		message := Message{"Hello r3x"}

		js, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":8080", nil)
}

