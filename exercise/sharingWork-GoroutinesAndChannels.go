package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("sharing work: Goroutines and Channels")

	go a()
	go b()
	fmt.Println("main end")
	time.Sleep(time.Second)

	// responseSizeofWeb("http://www.example.com/")
	// responseSizeofWeb("http://www.golang.org/")
	// responseSizeofWeb("http://www.golang.org/doc")
}

func responseSizeofWeb(website string) {
	response, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}
