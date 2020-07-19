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
	path      string
	faults    []reviewers.Fault
	reviewers []reviewers.Reviewer
}

func (r Runner) Run() error {
	config := LoadConfiguration()
	r.reviewers = config.SelectedReviewers()

	err := filepath.Walk(r.path, r.walkFn)

	if err != nil {
		return err
	}

	for _, fault := range r.faults {
		fmt.Println(fault)
	}

	if len(r.faults) > 0 {
		return ErrFaultsFound
	}

	return nil
}

func (r *Runner) walkFn(path string, info os.FileInfo, err error) error {
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

	for _, rev := range r.reviewers {
		fileFaults, err := rev.Review(path, reader)
		if err != nil {
			fmt.Printf("[Warning] Error executing %v : %v", rev.ReviewerName(), err)
			continue
		}

		r.faults = append(r.faults, fileFaults...)
	}

	return nil
}

func NewRunner(path string) *Runner {
	return &Runner{
		path: path,
	}
}
