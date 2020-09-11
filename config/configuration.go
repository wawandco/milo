// Configuration package is in charge of loading configuration from `.milo.yml` file.
package config

import (
	"fmt"
	"io/ioutil"
	"strings"

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

func Load() Settings {
	result := Settings{}

	data, err := ioutil.ReadFile(".milo.yml")
	if err != nil {
		fmt.Println("Running all reviewers, see more details in: https://github.com/wawandco/milo#configuration.")
		return result
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("[Warning] missformatted .milo.yml")
		return result
	}

	return result
}
