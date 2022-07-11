package main

import (
	"fmt"
)

func Run() error {
	fmt.Println("Starting out our application")
	return nil
}

func main() {
	fmt.Println("Simple Go rest api for answers")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
