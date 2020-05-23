package milo

import (
	"fmt"
	"os"
	"path/filepath"
	"wawandco/milo/reviewers"
)

type Runner struct {
	path string
}

func (r Runner) Run() error {

	referee := NewReferee()
	referee.Reviewers = []Reviewer{
		reviewers.Doctype{},
	}

	var faults []reviewers.Fault
	err := filepath.Walk(r.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		//Open the file
		reader, err := os.Open(path)
		if err != nil {
			return err
		}

		fileFaults, err := referee.Review(path, reader)
		if err != nil {
			return err
		}

		faults = append(faults, fileFaults...)
		return nil
	})

	fmt.Println(faults)

	return err
}

func NewRunner(path string) (*Runner, error) {
	runner := &Runner{
		path: path,
	}

	return runner, nil
}
