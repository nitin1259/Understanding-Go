package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
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

	sub := defaultSubscriber("Nitin")
	printSubscriber(sub)

}

func defaultSubscriber(name string) *subscriber {
	var sub subscriber
	sub.name = name
	sub.rate = 10.0
	sub.active = true

	return &sub
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
