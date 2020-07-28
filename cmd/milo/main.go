package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"text/tabwriter"

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

type CommandHelper interface {
	cmd.Runnable

	HelpText() string
}

func printHelp() {
	result := "Milo checks for issues with your HTML code.\n\n"
	result += "Usage:\n"
	result += "  milo [command] [args]\n\n"
	result += "Available Commands:"
	fmt.Print(result)

	// initialize tabwriter
	w := new(tabwriter.Writer)
	defer w.Flush()

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 3, '\t', 0)

	for _, runnable := range runnables {
		c, ok := runnable.(CommandHelper)
		if !ok {
			continue
		}

		fmt.Fprintf(w, "\n %v\t%v", c.Name(), c.HelpText())
	}
}
