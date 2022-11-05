package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/covrom/herthbot/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	a := app.New()
	a.Serve(ctx)
	log.Println("Exited.")
}
