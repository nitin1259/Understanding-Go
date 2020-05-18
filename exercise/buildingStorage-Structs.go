package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type Employee struct {
	name    string
	salary  float64
	homeAdd Address
}

type Address struct {
	street string
	city   string
	state  string
	zip    string
}

type part struct {
	description string
	count       int
}

func main() {
	fmt.Println("building storage: Structs")

	/*
		var mystruct struct {
			num    float64
			name   string
			toogle bool
		}

		fmt.Printf("%#v \n", mystruct)

		mystruct.name = "nitin"
		mystruct.num = 23.4
		mystruct.toogle = true

		fmt.Println(mystruct.name, mystruct)

		// defined types and structs

		type car struct {
			name     string
			topspeed int
		}

		var suzuki car

		suzuki.name = "maruti"
		suzuki.topspeed = 140

		fmt.Println("struct type suzuki: ", suzuki)

		var partInput part
		partInput.count = 20
		partInput.description = "part one"
		showInfo(partInput)

		p := minimumPartOrder("min order", 60)
		fmt.Println("part order value: ", p)

		updatePartOrder(&p)
		fmt.Println("updated part order: ", p)
	*/

	/*
		sub := defaultSubscriber("Nitin")
		applyDiscount(sub)
		printSubscriber(sub)

		// String literals
		sub2 := subscriber{
			name:   "kapil",
			rate:   9.99,
			active: false,
		}

		printSubscriber(&sub2)
	*/

	address := Address{street: "dalia lama", city: "Sanghai", state: "kajakode", zip: "34221"}
	fmt.Println("Address: ", address)

	emp := Employee{name: "Sachin", salary: 23456.7, homeAdd: address}
	fmt.Printf("Employee details: %#v \n", emp)

	fmt.Println(emp.homeAdd.state)
}

func defaultSubscriber(name string) *subscriber {
	var sub subscriber
	sub.name = name
	sub.rate = 10.0
	sub.active = true

	return &sub
}
func applyDiscount(sub *subscriber) {
	sub.rate = 5.0
}

func printSubscriber(sub *subscriber) {

	fmt.Println("name: ", sub.name)
	fmt.Println("rate: ", sub.rate)
	fmt.Println("active?: ", sub.active)
}

func showInfo(input part) {

	fmt.Printf("part description: %s and count: %d \n", input.description, input.count)
}

func minimumPartOrder(description string, count int) part {
	var partOrder part
	partOrder.description = description
	partOrder.count = count

	return partOrder
}

func updatePartOrder(input *part) {

	fmt.Println("printing updateorder part input: ", input)
	// *input.count=99
	//If you write *pointer.myField, Go thinks that myField must contain a pointer. It doesn’t, though, so an error results.
	// To get this to work, you need to wrap *pointer in parentheses.
	// That will cause the myStruct value to be retrieved, after which you can access the struct field.
	(*input).count = 99
	// or
	input.count = 99
	// return input
}
