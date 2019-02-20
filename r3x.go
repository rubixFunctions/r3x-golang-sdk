package r3x

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)


func Execute (r3xFunc func() []byte) {
	HTTPStream(r3xFunc)
}

func HTTPStream(r3xFunc func() []byte){
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		b := r3xFunc()

		var f interface{}

		err := json.Unmarshal(b, &f)
		if err != nil {
			fmt.Println("error:", err)
		}

		js, err := json.MarshalIndent(&f, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(js)
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}

