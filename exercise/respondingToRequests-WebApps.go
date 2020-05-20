package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("responding to requests: Web Apps")
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello from web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

// Go supports first-class function, functions can be assigned to variables, and then called from those variables.

func sayHi() {
	fmt.Println("say hi from first class function")
}
