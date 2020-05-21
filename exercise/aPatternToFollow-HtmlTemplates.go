package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("a pattern to follow: HTML Templates")

	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8090", nil)
	log.Fatal(err)

}

func viewHandler(writer http.ResponseWriter, req *http.Request) {

	resp := []byte("Welcome to guestbook")
	_, err := writer.Write(resp)
	if err != nil {
		log.Fatal(err)
	}
}
