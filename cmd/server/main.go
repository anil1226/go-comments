package main

import (
	"context"
	"fmt"

	"github.com/anil1226/go-banking/internal/db"
)

func main() {
	fmt.Println("main!!")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
	fmt.Println("Starting up our application!!")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to db")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected to db")
	return nil
}
