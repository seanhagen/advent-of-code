package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("neead at least the name of the script to run!\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		get()

	case "part2":
		part2()
	}

}
