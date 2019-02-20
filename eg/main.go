package main

import (
	"github.com/rubixFunctions/r3x-golang-sdk"
)

func main(){
	r3x.Execute(r3xFunc)
}

func r3xFunc() []byte {
	return []byte(`{"message": "hello r3x"}`)
}