package main

import (
	"context"
	"fmt"
	"github.com/JohnKucharsky/go_chi/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
