package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("on the list: Arrays")

	/*
		var primeArr [5]int // by default-  assing to Zero values of data type

		primeArr[0] = 2

		fmt.Println(primeArr[2])
		fmt.Println(primeArr[4])
		fmt.Println(primeArr[0])

		// array literals
		// var notes [7]string = [7]string{"sa", "re", "ga", "ma", "pa", "dha", "ne"}
		// or
		notes := [7]string{"sa", "re", "ga", "ma", "pa", "dha", "ne"} // short variable declaration
		primes := [5]int{2, 3, 5, 7, 11}
		fmt.Println(notes[4], notes[2], notes[6])
		fmt.Println(primes)

		// spread array literals over multiple lines
		text := [3]string{
			"this is first stirng literals",
			"this is second one",
			"only left is third one",
		}

		fmt.Println(text)

		fmt.Printf("%#v\n", text)

		// looping in array

		for i := 0; i < len(primes); i++ {
			fmt.Printf("index: %d, value: %d \n", i, primes[i])
		}

		for index, str := range notes {
			fmt.Printf("inex: %d, value: %s\n", index, str)
		}
	*/

	// reading a data from a file
	file, err := os.Open("./exercise/data.txt")
	if err != nil {
		log.Fatal("error while opening file: ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal("error while closing the file: ", err)
	}

	if scanner.Err() != nil {
		log.Fatal("error while scanning the file: ", scanner.Err())
	}

}
