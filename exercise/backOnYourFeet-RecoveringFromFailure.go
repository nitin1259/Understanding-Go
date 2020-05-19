package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("back on your feet: Recovering from Failure")

	// Deferring function calls
	// The defer keyword ensures a function call takes place even if the calling function exits early, say, by using the return keyword.

	err := doSocial()

	if err != nil {
		fmt.Println(err)
	}

	sl, err := readDirectory("keyboard")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sl)
}

func doSocial() error {
	defer fmt.Println("Goodbye")
	fmt.Println("Hi")
	return fmt.Errorf("I dont want to talk")
	fmt.Println("Hello howz you doing?")
	return nil
}

func readDirectory(dirName string) ([]string, error) {
	fileArr, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, fmt.Errorf("error while reading directory %s error: %s", dirName, err)
	}
	retSl := make([]string, 0)
	for _, file := range fileArr {
		var fileType string
		if fileType = "File"; file.IsDir() {
			fileType = "Directory"
		}
		// fmt.Println(file.IsDir(), file.Name())
		retSl = append(retSl, fmt.Sprintf("%s: %s", fileType, file.Name()))
	}
	return retSl, nil
}
