package main

import "fmt"

func main() {
	fmt.Println("main!!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
	fmt.Println("Starting up our application!!")
	return nil
}
