package main

import (
	"fmt"
	"log"

	"github.com/nitin1259/Understanding-Go/keyboard"
)

func main() {
	fmt.Println("Welcome to bundles of code : Packages")
	// If parts of your code are shared between multiple programs, you should consider moving them into packages.

	input, err := keyboard.GetFloadInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Entered input from standard Io: ", input)

	// Serving HTML documentation to yourself with “godoc”
	// The godoc tool generates HTML documentation based on the code in your main Go installation and your workspace.
	// To run godoc in web server mode, we’ll type the godoc command (again, don’t confuse that with go doc) in a terminal, followed by a special option: -http=:6060.
	// Then with godoc running, you can type the URL: http://localhost:6060/pkg
}
