package main

import (
	"bufio"
	"fmt"
	"log"
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