package main

import "fmt"

// you can use any type as an underlying type.

type Gallons float64
type Liters float64
type Milileters float64
type MyType string

func main() {

	fmt.Println("you’re my type: Defined Types")

	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240)

	fmt.Printf("Carfuel: %0.2f, Bus Fuel: %0.2f \n", carFuel, busFuel)

	// carFuel = Liters(20) not possible
	carFuel1 := Gallons(Liters(20)) // but you can convert between types that have the same underlying type
	busFuel1 := Liters(Gallons(250))
	fmt.Printf("Carfuel: %0.2f, Bus Fuel: %0.2f \n", carFuel1, busFuel1)

	carFuel1 = Gallons(Liters(40) * 0.264) // perform whatever operations are necessary to convert the underlying type value to a value appropriate for the type you’re converting to
	busFuel1 = Liters(Gallons(90) * 3.785) // A defined type supports all the same operations as its underlying type
	fmt.Printf("Carfuel: %0.2f, Bus Fuel: %0.2f \n", carFuel1, busFuel1)

	// define method on types
	someone := MyType("some one")
	someone.sayHi()

	anotherOne := MyType("another one") //Once a method is defined on a type, it can be called on any value of that type.
	anotherOne.sayHi()

	gal := Gallons(10)
	gal.doDouble()
	fmt.Println(gal)

	someLiters := Liters(50)
	litToGallons := someLiters.toGallons()
	fmt.Printf(" %0.3f litres = %0.3f gallons", someLiters, litToGallons)
}

func (mytype MyType) sayHi() {
	fmt.Printf("Say hi to %s !!!\n", mytype)
}

func (g *Gallons) doDouble() {
	*g *= 2
}

func (l *Liters) toGallons() Gallons {
	return Gallons(*l * 0.264)
}

func (m *Milileters) toGallons() Gallons {
	return Gallons(*m * 0.264 * 0.001)
}
