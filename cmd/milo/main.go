package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println("please pass the folder to analize, p.e: milo templates")
		return
	}

	runner := NewRunner(os.Args[1])

	err := runner.Run()
	if err != nil {
		log.Fatal(err)
	}
}
