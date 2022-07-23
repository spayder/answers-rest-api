package main

import (
	"fmt"
	"github.com/spayder/answers-rest-api/internal/answer"
	"github.com/spayder/answers-rest-api/internal/db"
	transportHttp "github.com/spayder/answers-rest-api/internal/transport/http"
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
	httpHandler := transportHttp.NewHandler(answerService)
	if err := httpHandler.Serve(); err != nil {
		fmt.Println("failed starting the handler")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Simple Go rest api for answers")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
