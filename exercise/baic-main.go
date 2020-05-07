package main

import (
	"fmt"
	"reflect"
)

func main() {
	var quantity int = 4

	// to check the type of variable
	fmt.Println(reflect.TypeOf(quantity))

	// shortHand variable declerations
	quant := 5
	len, wid := 20.5, 31.0
	name := "Nitin"

	fmt.Println(quant)
	fmt.Println(len, wid)
	fmt.Println(reflect.TypeOf(wid))
	fmt.Println(name)

	//conversion

	length := 16.2
	width := 2
	if length > float64(width) {
		fmt.Println("conversion of same type only possible")
	}
	fmt.Println(int32(length))

}
