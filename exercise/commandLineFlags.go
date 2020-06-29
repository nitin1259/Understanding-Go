package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Command-Line Flags")

	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 42, "some number")
	boolPtr := flag.Bool("isAvailable", false, "some boolean flag")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string variable")

	flag.Parse()

	fmt.Println("workPtr: ", *wordPtr)
	fmt.Println("numPtr: ", *numPtr)
	fmt.Println("boolPtr: ", *boolPtr)

	fmt.Println("svar: ", svar)

	fmt.Println("tail: ", flag.Args())
}
