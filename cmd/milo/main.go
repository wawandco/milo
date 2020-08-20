// Main milo command.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/help"
	"github.com/wawandco/milo/cmd/initialize"
	"github.com/wawandco/milo/cmd/review"
	"github.com/wawandco/milo/cmd/version"
)

var (
	// commands holds the list of commands that milo makes available
	// through the cli.
	commands = []cmd.Runner{
		review.Runner{},
		initialize.Runner{},
		version.Printer{},
	}

	printHelp = help.Printer{
		Commands: commands,
	}
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
		printHelp.Run(os.Args[1:])

		return
	}

	for _, command := range commands {
		if command.Name() != os.Args[1] {
			continue
		}

		err := command.Run(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	printHelp.Run(os.Args[1:])
}
