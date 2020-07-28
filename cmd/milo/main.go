package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/initialize"
	"github.com/wawandco/milo/cmd/review"
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
	result := "Milo checks for issues with your HTML code.\n\n"
	result += "Usage:\n"
	result += "  milo [command] [args]\n\n"
	result += "Available Commands:\n"

	for _, runnable := range runnables {
		c, ok := runnable.(CommandHelper)
		if !ok {
			continue
		}

		result += fmt.Sprintf("  %v\t%v\n", c.Name(), c.HelpText())
	}

	fmt.Println(result)
}
