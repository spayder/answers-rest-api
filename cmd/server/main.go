package main

import (
	"context"
	"fmt"
	"github.com/spayder/answers-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting out our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err = db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged to the database")
	return nil
}

func main() {
	fmt.Println("Simple Go rest api for answers")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
