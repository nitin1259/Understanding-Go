package main

import (
	"fmt"
	"log"

	"github.com/nitin1259/Understanding-Go/utils"
)

func main() {
	fmt.Println("keep it to yourself: Encapsulation and Embedding")

	// someDate := utils.Date{Year: 1989, Month: 12, Day: 31}
	// fmt.Println(someDate)

	// otherDate := Date{Year: 0, Month: 0, Day: -2}
	// fmt.Println(otherDate) // this is invalid date but till now we dont have control to set limitations

	anotherDate := utils.Date{}
	err := anotherDate.SetDay(30)
	err = anotherDate.SetMonth(11)
	err = anotherDate.SetYear(2019)
	if err != nil {
		log.Fatal("Entered : ", err)
	}
	fmt.Println(anotherDate)

	event := utils.Event{}
	event.SetTitle("moms birthday")
	// event.Date.Day ==23 // this will not work as its unexported in Date type
	errD := event.Date.SetDay(23)
	if errD != nil {
		log.Fatal("Entered : ", errD)
	}
	errM := event.SetMonth(10)
	if errM != nil {
		log.Fatal("Entered : ", errM)
	}
	errY := event.Date.SetYear(1981)
	if errY != nil {
		log.Fatal("Entered : ", errY)
	}

	fmt.Println("events: ", event)

	// errI := event.Date.SetYear(-1981)
	// if errI != nil {
	// 	log.Fatal("Entered : ", errI)
	// }
}
