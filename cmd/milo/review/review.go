// Package review is in charge of the review command. It will be invoked to review
// A set of files and print the faults found in that file.
package review

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/output"
	"github.com/wawandco/milo/reviewers"

	flag "github.com/spf13/pflag"
)

var (
	ErrFaultsFound      = errors.New("faults found")
	ErrInsufficientArgs = errors.New("please pass the folder to analyze, p.e: milo review file.html folder")
	ErrUnknownFormatter = errors.New("unknown formatter")
)

func NewRunner() *Runner {
	return &Runner{
		stdout: os.Stdout,
	}
}

// Runner is in charge of running the reviewers
// across the files passed.
type Runner struct {
	faults    []reviewers.Fault
	reviewers []reviewers.Reviewer
	formatter output.FaultFormatter

	stdout io.Writer
}

func (r Runner) Name() string {
	return "review"
}

func (r Runner) HelpText() string {
	return "looks for faults in files/folders passed as args."
}

func (r *Runner) SetOutput(w io.Writer) {
	r.stdout = w
}

func (r Runner) Run(args []string) error {
	flag.Parse()
	if len(args) < 2 {
		return ErrInsufficientArgs
	}

	c, err := config.Load()
	if errors.Is(err, config.ErrConfigFormat) {
		fmt.Fprintln(r.stdout, "[Warning] .milo.yml file is not in the correct format, using default values.")
	}

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

	for _, path := range args[1:] {
		err = filepath.Walk(path, r.walkFn)
		if err != nil {
			return err
		}
	}

	for _, fault := range r.faults {
		ou := r.formatter.Format(fault)
		if ou == "" {
			continue
		}

		fmt.Fprintln(r.stdout, ou)
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

	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if filepath.Ext(path) != ".html" {
		return nil
	}

	for _, rev := range r.reviewers {
		fileFaults, err := rev.Review(path, bytes.NewBuffer(data))
		if err != nil {
			fmt.Fprintf(r.stdout, "[Error] Error executing %v : %v", rev.ReviewerName(), err)

			continue
		}

		r.faults = append(r.faults, fileFaults...)
	}

	return nil
}
