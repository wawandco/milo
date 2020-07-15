package runtime

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"wawandco/milo/reviewers"
)

var (
	ErrFaultsFound = errors.New("faults found")
)

// Runner is in charge of initializing a Referee with
// the reviewers we have in the app.
type Runner struct {
	path string
}

func (r Runner) Run() error {
	config := LoadConfiguration()
	revs := config.SelectedReviewers()

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

		for _, rev := range revs {
			fileFaults, err := rev.Review(path, reader)
			if err != nil {
				fmt.Printf("[Warning] Error executing %v : %v", rev.ReviewerName(), err)
				continue
			}

			faults = append(faults, fileFaults...)
		}

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
