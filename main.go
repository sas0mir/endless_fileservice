package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/sas0mir/endless_fileservice/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Error starting the application:", err)
	}
}
