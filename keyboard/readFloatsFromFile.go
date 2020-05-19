package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// GetFloatsFromFile function will read the inputs from file
// return the float array and error
func GetFloatsFromFile(filePath string) ([3]float64, error) {

	var numbers [3]float64
	i := 0
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return numbers, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Println(err)
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return numbers, err
	}

	if scanner.Err() != nil {
		log.Print(scanner.Err())
		return numbers, scanner.Err()
	}
	return numbers, nil
}
