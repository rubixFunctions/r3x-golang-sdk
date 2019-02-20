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

		if r.Method != "POST" {
			http.Error(w, "Invalid Request", http.StatusInternalServerError)
			return
		}

		m := jsonHandler(w, r)

		b := r3xFunc(m)

		var f interface{}

		err := json.Unmarshal(b, &f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		js, err := json.MarshalIndent(&f, "", "\t")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(js)
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}

func jsonHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	body, err := ioutil.ReadAll(r.Body)
	defer  r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var m map[string]interface{}

	if len(body) > 0 {
		var bf interface{}

		err = json.Unmarshal(body, &bf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("error:", err)
		}

		m = bf.(map[string]interface{})
	}

	return m
}

