package r3x

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func Execute (r3xFunc func(map[string]interface{}) []byte) {
	HTTPStream(r3xFunc)
}

func HTTPStream(r3xFunc func(map[string]interface{}) []byte){
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		body, err := ioutil.ReadAll(r.Body)
		defer  r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var m map[string]interface{}

		if len(body) > 0 {
			var bf interface{}

			err = json.Unmarshal(body, &bf)
			if err != nil {
				fmt.Println("error:", err)
			}

			m = bf.(map[string]interface{})
		}

		b := r3xFunc(m)

		var f interface{}

		err = json.Unmarshal(b, &f)
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

