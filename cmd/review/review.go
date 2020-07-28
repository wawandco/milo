package review

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/output"
	"github.com/wawandco/milo/reviewers"
)

var (
	ErrFaultsFound      = errors.New("faults found")
	ErrInsufficientArgs = errors.New("please pass the folder to analize, p.e: milo run templates")
)

// Runner is in charge of running the reviewers
// across the files passed.
type Runner struct {
	faults    []reviewers.Fault
	reviewers []reviewers.Reviewer
	formatter output.FaultFormatter
}

func (r Runner) Name() string {
	return "run"
}

func (r Runner) Run(args []string) error {
	if len(args) < 2 {
		return ErrInsufficientArgs
	}

	c := config.LoadConfiguration()
	r.reviewers = c.SelectedReviewers()
	r.formatter = c.Printer()

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

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	for _, rev := range r.reviewers {
		fileFaults, err := rev.Review(path, bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("[Warning] Error executing %v : %v", rev.ReviewerName(), err)
			continue
		}

		r.faults = append(r.faults, fileFaults...)
	}

	return nil
}
