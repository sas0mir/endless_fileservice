package main

import (
	"context"
	"fmt"

	"github.com/sas0mir/endless_fileservice/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Error starting the application:", err)
	}
}
