package main

import "fmt"

var version = "latest"

type VersionPrinter struct{}

func (v VersionPrinter) Name() string {
	return "version"
}

func (v VersionPrinter) Run([]string) error {
	fmt.Printf("Running Milo %v\n", version)
	return nil
}

func (v VersionPrinter) HelpText() string {
	return "prints the Milo version"
}
