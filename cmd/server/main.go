package main

import (
	"context"
	"fmt"
	"github.com/spayder/answers-rest-api/internal/answer"
	"github.com/spayder/answers-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting out our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.Migrate(); err != nil {
		fmt.Println("failed to migrate the database")
		return err
	}

	fmt.Println("successfully connected and pinged to the database")

	answerService := answer.NewService(db)
	fmt.Println(answerService.GetAnswer(
		context.Background(),
		"7210a7da-0372-11ed-b939-0242ac120002",
	))

	return nil
}

func main() {
	fmt.Println("Simple Go rest api for answers")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
