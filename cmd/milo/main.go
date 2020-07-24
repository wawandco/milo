package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"wawandco/milo/cmd"
	"wawandco/milo/cmd/initialize"
	"wawandco/milo/cmd/review"
)

var runnables = []cmd.Runnable{
	review.Runner{},
	initialize.Runner{},
}

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
		printHelp()
		return
	}

	for _, runnable := range runnables {
		if runnable.Name() != os.Args[1] {
			continue
		}

		err := runnable.Run(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	printHelp()
}

func printHelp() {
	fmt.Println(`[here goes the help text]`)
}
