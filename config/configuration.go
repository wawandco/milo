// Configuration package is in charge of loading configuration from `.milo.yml` file.
package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/wawandco/milo/output"
	"github.com/wawandco/milo/reviewers"

	"gopkg.in/yaml.v2"
)

// Settings will load the linting configuration from .milo.yml.
type Settings struct {
	Output    string
	Reviewers []string
}

func (c Settings) SelectedReviewers() []reviewers.Reviewer {
	if len(c.Reviewers) == 0 {
		return reviewers.All
	}

	selected := []reviewers.Reviewer{}
	for _, reviewer := range reviewers.All {
		if !strings.Contains(strings.Join(c.Reviewers, " || "), reviewer.ReviewerName()) {
			continue
		}

		selected = append(selected, reviewer)
	}

	return selected
}

func (c Settings) Printer() output.FaultFormatter {
	for _, printer := range output.Formatters {
		if printer.FormatterName() != c.Output {
			continue
		}

		return printer
	}

	return output.TextFaultFormatter{}
}

func LoadConfiguration() Settings {
	result := Settings{}

	data, err := ioutil.ReadFile(".milo.yml")
	if err != nil {
		fmt.Println("[Warning] could not load configuration from .milo.yml")
		return result
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("[Warning] missformatted .milo.yml")
		return result
	}

	return result
}
