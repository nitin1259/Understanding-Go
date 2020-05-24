package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	fmt.Println("back on your feet: Recovering from Failure")

	// panic("oh no some wired happens")

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

	scanDirectory("keyboard")

	defer afterScanDirectory()
	panicCreate()
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

// recursive function to scan a directory for all files
func scanDirectory(dirName string) error {
	fmt.Println(dirName)
	fileArr, err := ioutil.ReadDir(dirName)
	if err != nil {
		return fmt.Errorf("error while reading directory %s error: %s", dirName, err)
	}
	// retSl := make([]string, 0)
	for _, file := range fileArr {
		filepath := filepath.Join(dirName, file.Name())
		if file.IsDir() {
			err = scanDirectory(filepath)
			if err != nil {
				return fmt.Errorf("error while reading directory %s error: %s", dirName, err)
			}
		} else {
			fmt.Println(filepath)
		}

		// fmt.Println(file.IsDir(), file.Name())
	}
	return nil
}

// Deferred calls completed before crash

// recursive function to scan a directory for all files
func scanDirectoryWithpanic(dirName string) {
	fmt.Println(dirName)
	fileArr, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	// retSl := make([]string, 0)
	for _, file := range fileArr {
		filepath := filepath.Join(dirName, file.Name())
		if file.IsDir() {
			err = scanDirectory(filepath)
			if err != nil {
				fmt.Println("error while reading directory %s error: %s", dirName, err)
			}
		} else {
			fmt.Println(filepath)
		}
	}
}

func panicCreate() {
	fmt.Println("Before panic")
	panic(fmt.Errorf("Create err or from panicCreate"))
	fmt.Println("after panic")
}

func afterScanDirectory() {
	// Calling recover will not cause execution to resume at the point of the panic, at least not exactly.
	// The function that panicked will return immediately, and none of the code in that function’s block following
	// the panic will be executed. After the function that panicked returns, however, normal execution resumes.
	p := recover()
	if p == nil {
		fmt.Println("no panic happens")
	}

	err, ok := p.(error)
	if ok {
		fmt.Println("happen becuase of err: ", err)
	}
}
