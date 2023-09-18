package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
	} else {
		z, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Print(err)
		}
		// Process the file data here
		fmt.Print(string(z))
	}
}
