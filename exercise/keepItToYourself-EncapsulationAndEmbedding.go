package main

import (
	"errors"
	"fmt"
	"log"
)

// Date struct type
type Date struct {
	Year  int
	Month int
	Day   int
}

// SetYear setter methods
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

// SetMonth setter method
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

// SetDay setter method
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	fmt.Println("keep it to yourself: Encapsulation and Embedding")

	someDate := Date{Year: 1989, Month: 12, Day: 31}
	fmt.Println(someDate)

	otherDate := Date{Year: 0, Month: 0, Day: -2}
	fmt.Println(otherDate) // this is invalid date but till now we dont have control to set limitations

	anotherDate := Date{}
	err := anotherDate.SetDay(30)
	err = anotherDate.SetMonth(-11)
	err = anotherDate.SetYear(2019)
	if err != nil {
		log.Fatal("Entered : ", err)
	}
	fmt.Println(anotherDate)

}
