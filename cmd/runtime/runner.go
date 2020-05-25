package runtime

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"wawandco/milo"
	"wawandco/milo/reviewers"
)

var (
	ErrFaultsFound = errors.New("faults found")
)

type Runner struct {
	path string
}

func (r Runner) Run() error {
	referee := milo.NewReferee()
	referee.Reviewers = []milo.Reviewer{
		reviewers.DoctypePresent{},
		reviewers.DoctypeValid{},
		reviewers.InlineCSS{},
		reviewers.TitlePresent{},
		reviewers.StyleTag{},
		reviewers.TagLowercase{},
	}

	var faults []reviewers.Fault
	err := filepath.Walk(r.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

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

	if err != nil {
		return err
	}

	for _, fault := range faults {
		fmt.Println(fault)
	}

	if len(faults) > 0 {
		return ErrFaultsFound
	}

	return nil
}

func NewRunner(path string) *Runner {
	return &Runner{
		path: path,
	}
}
