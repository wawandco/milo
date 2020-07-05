package runtime

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"wawandco/milo"
	"wawandco/milo/reviewers"

	"gopkg.in/yaml.v2"
)

var (
	ErrFaultsFound = errors.New("faults found")
)

type Runner struct {
	path string
}

type Configuration struct {
	Enabled []string
}

func (r Runner) Run() error {
	yamlFile, errReadFile := ioutil.ReadFile("./conf.yml")
	if errReadFile != nil {
		return errReadFile
	}

	conf := Configuration{}

	errUnmarshal := yaml.Unmarshal(yamlFile, &conf)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	referee := milo.NewReferee()
	referee.SetEnabledReviewers(conf.Enabled)

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
