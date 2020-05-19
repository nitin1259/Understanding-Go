package main

import (
	"fmt"
	"reflect"
)

// MyType1 type with underlying int type
type MyType1 int

// MyType2 type with underlying int type
type MyType2 int

// MyInterface for example
// In Go, an interface is defined as a set of methods that certain values are expected to have.
// You can think of an interface as a set of actions you need a type to be able to perform
type MyInterface interface {
	MethodWithOutParamter()
	MethodWithParameter(float64)
	MethodWithReturnType() string
}

// MethodWithOutParamter interface
func (m MyType1) MethodWithOutParamter() {
	fmt.Println("Method with out parameter", m)
}

// MethodWithParameter interface method
func (m MyType1) MethodWithParameter(input float64) {
	fmt.Println("Method with parameter: ", input)

}

//MethodWithReturnType interface method
func (m MyType1) MethodWithReturnType() string {
	fmt.Println("Method with return type")
	return "Method with Return type"
}

// MethodNotInInterface not from interface
func (m *MyType1) MethodNotInInterface() {
	fmt.Println("Method not in interface")
}

// MethodWithOutParamter interface
func (m MyType2) MethodWithOutParamter() {
	fmt.Println("2 Method with out parameter", m)
}

// MethodWithParameter interface method
func (m MyType2) MethodWithParameter(input float64) {
	fmt.Println("2 Method with parameter: ", input)

}

//MethodWithReturnType interface method
func (m MyType2) MethodWithReturnType() string {
	fmt.Println(" 2Method with return type")
	return " 2 Method with Return type"
}

// MethodNotInInterface not from interface
func (m *MyType2) MethodNotInInterface() {
	fmt.Println("2 Method not in interface")
}

func main() {
	fmt.Println("what can you do?: Interfaces")

	var value1 MyInterface
	value1 = MyType1(1)

	value1.MethodWithOutParamter()
	value1.MethodWithReturnType()
	fmt.Println(reflect.TypeOf(value1))

	value2 := MyType2(4)
	fmt.Println(value2.MethodWithReturnType())

	genericMyInterfaceParameter(value1)
	genericMyInterfaceParameter(value2)

	//You can only call methods defined as part of the interface
	// value1.MethodNotInInterface() // this will produce error.
	value2.MethodNotInInterface()

	// Type assertions
	tryOurNonInterfaceMethod(value1)

	tryOurNonInterfaceMethod(value2)

}

func genericMyInterfaceParameter(inputInf MyInterface) {

	inputInf.MethodWithOutParamter()
	inputInf.MethodWithReturnType()
	// inputInf.MethodNotInInterface() // this will produce error.
}

func tryOurNonInterfaceMethod(inputInf MyInterface) {
	// inputInf.MethodNotInInterface() // this method is not part of interface

	// When you have a value of a concrete type assigned to a variable with an interface type,
	// a type assertion lets you get the concrete type back. It’s kind of like a type conversion.
	//  Its syntax even looks like a cross between a method call and a type conversion.
	//  After an interface value, you type a dot, followed by a pair of parentheses with the concrete type.
	// (Or rather, what you’re asserting the value’s concrete type is.)
	mytype1, ok := inputInf.(MyType1)
	if ok {

		mytype1.MethodNotInInterface()
	}

	if !ok {
		fmt.Println("assertion type is not succesful")
	}

}
