package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	// "reflect"
	"strconv"
	"strings"
	
)

func main() {
	fmt.Println("In condition and loops ...")

	/*
	var now time.Time = time.Now()
	year := now.Year()
	fmt.Println(year)

	// another example of taking input from constole
	fmt.Print("Enter the percentage: ")
	reader:= bufio.NewReader(os.Stdin)
	percentage, err:= reader.ReadString('\n');
	if err!=nil{
		log.Fatalf("error while reading input %s", err)
	}

	// blank identifier to ignore the error	
	// percentage, _:= reader.ReadString('\n');
	fmt.Printf("entered percentage: %s", reflect.TypeOf(percentage) )
	fmt.Println()

	input:= strings.TrimSpace(percentage);

	per, errFlt := strconv.ParseFloat(input, 64)

	if errFlt!=nil{
		log.Fatalf("Entered nmber is not correct %s", errFlt)
	}

	if(per>=60){
		fmt.Println("Pass")
	}else{
		fmt.Println("Fail")
	}

	// exapmle - get the file size

	fileInfo, errFile := os.Stat("./README.md");
	if errFile!=nil{
		log.Fatalf("Error while reading the file %s ", errFile)
	}

	size:= fileInfo.Size();
	fmt.Print("size: ");fmt.Println(size);

	*/

	// new game stated

	seconds:= time.Now().Unix()
	rand.Seed(seconds)
	target:= rand.Intn(100)+1; //problem here is it will same number every time
	fmt.Println(target)

	fmt.Println("We have choosen a random number between 1 and 100.")
	fmt.Println("Can you guess it ???")

	reader := bufio.NewReader(os.Stdin)
	success:=false;

	for guesses:= 0; guesses < 5; guesses++ {
		fmt.Println("Total number of guess left: ", 10-guesses)
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
	
		if err!=nil{
			log.Fatalf("error while reading input: %s", err)
		}
	
		input = strings.TrimSpace(input)
	
		inputNum, err := strconv.Atoi(input)
	
		if err !=nil{
			log.Fatalf("Enter number is not correct: %s", err)
		}
	
		if inputNum == target{
			success=true;
			fmt.Println("Congratulation! You guessed it correct")
			break;
		} else if inputNum > target{
			fmt.Println("Oops. You guess was HIGH.")
		}else{
			fmt.Println("Oops. You guess was LOW.")
		}
	}
	

	if !success{
		fmt.Println("Sorry! you are not able to guess. Target was: ", target)
	}
}