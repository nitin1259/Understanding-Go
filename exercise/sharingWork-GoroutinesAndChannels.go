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

	// go a()
	// go b()
	// fmt.Println("main end")
	// time.Sleep(time.Second)

	sizeChan := make(chan int)
	go responseSizeofWeb("http://www.example.com/", sizeChan)
	fmt.Println("size: ", <-sizeChan)
	go responseSizeofWeb("http://www.golang.org/", sizeChan)
	fmt.Println("size: ", <-sizeChan)
	go responseSizeofWeb("http://www.golang.org/doc", sizeChan)
	fmt.Println("size: ", <-sizeChan)
	// time.Sleep(time.Second * 5)
	/*
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

		// another example
		channel1 := make(chan string)
		channel2 := make(chan string)
		go abc(channel1)
		go def(channel2)

		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
		fmt.Println(<-channel1)
		fmt.Println(<-channel2)
	*/

	// my_channel := make(chan string)
	// go send(my_channel)
	// reportNap("receving go-routine", 5)
	// fmt.Println(<-my_channel)
	// fmt.Println(<-my_channel)

}

func reportNap(name string, sec int) {
	for i := 0; i < sec; i++ {
		fmt.Println(name, "Sleeping..")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up")
}

func send(myChannel chan string) {
	reportNap("sending go-routine", 2)
	fmt.Println("******** Sending value *******")
	myChannel <- "a"
	fmt.Println("******** Sending value *******")
	myChannel <- "b"

}

func abc(chan1 chan string) {
	chan1 <- "a"
	chan1 <- "b"
	chan1 <- "c"
}

func def(chan2 chan string) {
	chan2 <- "d"
	chan2 <- "e"
	chan2 <- "f"
}

func responseSizeofWeb(url string, sizeChan chan int) {
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
	sizeChan <- len(body)
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
