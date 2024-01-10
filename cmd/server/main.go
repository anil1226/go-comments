package main

import (
	"fmt"

	"github.com/anil1226/go-banking/internal/comment"
	"github.com/anil1226/go-banking/internal/db"
	transportHttp "github.com/anil1226/go-banking/internal/transport/http"
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

	fmt.Println("Successfully connected to db")

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	// cmtService.PostComment(
	// 	context.Background(),
	// 	comment.Comment{
	// 		ID:     "40e6215d-b5c6-4896-987c-f30f3678f608",
	// 		Slug:   "test",
	// 		Author: "Anil",
	// 		Body:   "Hi!!",
	// 	},
	// )

	// fmt.Println(cmtService.GetComment(context.Background(), "40e6215d-b5c6-4896-987c-f30f3678f608"))

	return nil
}
