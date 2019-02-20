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
	var response string
	if name != nil {
		response = fmt.Sprintf(`{"message": "hello %s"}`, name)
	} else {
		response = `{"message": "no input r3x"}`
	}
	return []byte(response)
}