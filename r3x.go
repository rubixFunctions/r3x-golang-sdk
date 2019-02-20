package r3x

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func Execute (r3xFunc func() []byte) {
	HTTPStream(r3xFunc)
}

func HTTPStream(r3xFunc func() []byte){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		fmt.Println(r3xFunc())

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

	http.ListenAndServe(":8080", nil)
}

