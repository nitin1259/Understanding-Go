package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("appending issue: Slices")
	// slices
	/*
	// basic of slice
	var primes []int
	fmt.Printf("primes: %#v len: %d \n",primes, len(primes))
	primes = make([]int, 5)
	primes[0]=2
	notes:= make([]string, 7);
	notes[4]= "pa"

	letters:= []string{"a", "b", "c"}
	for index, let:= range letters{
		fmt.Println(index, let)
	}
	*/

	/*
	numbers, err:= getFloatsFromFile()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(numbers)
	sum:=0.0
	for _,num:=range numbers{
		sum += num
	}

	fmt.Printf("Average value: %0.2f \n", sum/float64(len(numbers)))
	*/
	// Getting command-line arguments from the os.Args slice

	fmt.Println(os.Args) // go run exercise/appendingIssue-Slices.go 45.5 56.3 67.2 23.6
	cmdArgs:= os.Args;
	cmdArgs = cmdArgs[1:]
	fmt.Println(cmdArgs)


	// Variadic functions
	// nonvariadic arguments are always required; it’s a compile error to omit those. 
	// Only the last parameter in a function definition can be variadic; you can’t place it in front of required parameters

	fmt.Println(findMaxFloat(23.4, 29.1, 5.32, 67.3))

	someNumbers:= []float64{23.4, 29.1, 5.32, 67.3}
	fmt.Println(findMaxFloat(someNumbers...))// way of passing a same type slice to variacdic funciton...

}

func findMaxFloat(numbers ...float64)float64{
	max:= math.Inf(-1);
	for _, num:= range numbers{
		if num>max{
			max = num
		}
	}
	return max;
}

func getFloatsFromFile() ([]float64, error){
	var numbers []float64
	i := 0
	file, err := os.Open("./exercise/data.txt")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		num, err:= strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		numbers = append(numbers, num)
		i++
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if scanner.Err() != nil {
		log.Print(scanner.Err())
		return nil, scanner.Err()
	}
	return numbers, nil;
}