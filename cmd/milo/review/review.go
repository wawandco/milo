// Package review is in charge of the review command. It will be invoked to review
// A set of files and print the faults found in that file.
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

	flag "github.com/spf13/pflag"
)

var (
	ErrFaultsFound      = errors.New("faults found")
	ErrInsufficientArgs = errors.New("please pass the folder to analyze, p.e: milo review templates")
	ErrUnknownFormatter = errors.New("unknown formatter")
)

// Runner is in charge of running the reviewers
// across the files passed.
type Runner struct {
	faults    []reviewers.Fault
	reviewers []reviewers.Reviewer
	formatter output.FaultFormatter
}

func (r Runner) Name() string {
	return "review"
}

func (r Runner) HelpText() string {
	return "looks for faults in files/folder passed in the first arg."
}

func (r Runner) Run(args []string) error {
	flag.Parse()

	if len(args) < 2 {
		return ErrInsufficientArgs
	}

	c := config.LoadConfiguration()
	r.reviewers = c.SelectedReviewers()
	r.formatter = output.Formatter(c.Output)

	if r.formatter == nil {
		r.formatter = output.TextFaultFormatter{}
	}

	if cliOutput != "" {
		r.formatter = output.Formatter(cliOutput)
		if r.formatter == nil {
			return ErrUnknownFormatter
		}
	}

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
		return fmt.Errorf("%d %w", len(r.faults), ErrFaultsFound)
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

	if filepath.Ext(path) != ".html" {
		return nil
	}

	for _, rev := range r.reviewers {
		fileFaults, err := rev.Review(path, bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("[Error] Error executing %v : %v", rev.ReviewerName(), err)

			continue
		}

		r.faults = append(r.faults, fileFaults...)
	}

	return nil
}
