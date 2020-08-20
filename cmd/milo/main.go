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
	// runners holds the list of runners that milo makes available
	// through the cli.
	runners = []cmd.Runner{
		review.Runner{},
		initialize.Runner{},
		version.Printer{},
	}

	printHelp = help.Printer{
		Runners: runners,
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
		err := printHelp.Run(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	for _, runner := range runners {
		if runner.Name() != os.Args[1] {
			continue
		}

		err := runner.Run(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	err := printHelp.Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
