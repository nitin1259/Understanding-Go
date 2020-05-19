package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"
)

func main() {
	fmt.Println("sharing work: Goroutines and Channels")

	// go a()
	// go b()
	// fmt.Println("main end")
	// time.Sleep(time.Second)

	go responseSizeofWeb("http://www.example.com/")
	go responseSizeofWeb("http://www.golang.org/")
	go responseSizeofWeb("http://www.golang.org/doc")
	time.Sleep(time.Second * 5)

	// channels
	var myChannel chan float64
	myChannel = make(chan float64)
	// myChannel <- 3.14 // sending value to channel
	// <-myChannel       // receiving value from channel
	fmt.Println(reflect.TypeOf(myChannel), myChannel)

	someChannel := make(chan string)

	//Synchronizing goroutines with channels
	go greetings(someChannel)
	reciever := <-someChannel
	fmt.Println(reciever)
}

func responseSizeofWeb(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
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

func greetings(myChan chan string) {
	myChan <- "Greetng from channels"
}
