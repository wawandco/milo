package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"wawandco/milo/cmd/runtime"
)

type command interface {
	Run([]string) error
	CommandName() string
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

	commands := []command{
		runtime.Runner{},
		// init.Command{},
	}

	for _, command := range commands {
		if command.CommandName() != os.Args[1] {
			continue
		}

		err := command.Run(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		break
	}

	printHelp()
}

func printHelp() {
	fmt.Println(`[here goes the help text]`)
}
