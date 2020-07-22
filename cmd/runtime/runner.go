package runtime

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"wawandco/milo/output"
	"wawandco/milo/reviewers"
)

var (
	ErrFaultsFound = errors.New("faults found")
)

// Runner is in charge of initializing a Referee with
// the reviewers we have in the app.
type Runner struct {
	faults    []reviewers.Fault
	reviewers []reviewers.Reviewer
	formatter output.FaultFormatter
}

func (r Runner) CommandName() string {
	return "run"
}

func (r Runner) Run(args []string) error {
	if len(args) < 2 {
		return errors.New("please pass the folder to analize, p.e: milo run templates")
	}

	config := LoadConfiguration()
	r.reviewers = config.SelectedReviewers()
	r.formatter = config.Printer()

	err := filepath.Walk(args[1], r.walkFn)
	if err != nil {
		return err
	}

	for _, fault := range r.faults {
		ou := r.formatter.Format(fault)
		if ou == "" {
			continue
		}

		fmt.Println(ou)
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
