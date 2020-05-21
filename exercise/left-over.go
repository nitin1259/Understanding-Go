package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("left over : 6 things")

	fmt.Println(os.FileMode(0700))

	// #1 Initialization statements for “if”
	if err := initializeOnlyforIf(5); err != nil {
		log.Fatal(err)
	}

	if err := initializeOnlyforIf(-1); err != nil {
		log.Fatal(err)
	}

	// #2 The switch statement
	// Go automatically exits the switch at the end of a case’s code. no break required in every case.
	// There’s a fallthrough keyword you can use in a case, if you do want the next case’s code to run as well.

}

func initializeOnlyforIf(num int) error {
	// do some stuff
	if num < 0 {
		return fmt.Errorf("Some error occurs here: %d", num)
	}
	return nil
}
