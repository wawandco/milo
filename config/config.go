// Configuration package is in charge of loading configuration from `.milo.yml` file.
package config

import (
	"errors"
	"os"
	"strings"

	"github.com/wawandco/milo/reviewers"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigNotFound = errors.New("milo.yml not found")
	ErrConfigFormat   = errors.New("missformatted .milo.yml")
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

func Load() (Settings, error) {
	result := Settings{}

	data, err := os.ReadFile(".milo.yml")
	if err != nil {
		return result, ErrConfigNotFound
	}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return result, ErrConfigFormat
	}

	return result, nil
}
