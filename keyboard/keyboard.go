// Package keyboard is for taking input from keyboard
package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFloadInput function is to take the input form the standard input
// return the float64 and error
func GetFloadInput() (float64, error) {
	fmt.Print("Enter the number: ")

	reader := bufio.NewReader(os.Stdin)

	inputStr, err := reader.ReadString('\n')

	inputStr = strings.TrimSpace(inputStr)
	if err != nil {
		log.Println("Error while reading input: ", err)
		return 0, err
	}

	inputFloat, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		log.Print("Error while pasrsing the input to FLoat", err)
		return 0, err
	}

	return inputFloat, nil
}
