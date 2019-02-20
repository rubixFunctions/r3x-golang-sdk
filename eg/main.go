package main

import (
	"fmt"
	"github.com/rubixFunctions/r3x-golang-sdk"
)

func main(){
	r3x.Execute(r3xFunc)
}

func r3xFunc(input map[string]interface{}) []byte {
	name := input["name"]
	response := fmt.Sprintf(`{"message": "hello %s"}`, name)
	return []byte(response)
}