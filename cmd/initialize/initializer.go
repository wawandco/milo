package initialize

import (
	"io/ioutil"
	"wawandco/milo/config"
	"wawandco/milo/reviewers"

	"gopkg.in/yaml.v2"
)

type Runner struct{}

func (r Runner) Name() string {
	return "init"
}

func (r Runner) Run(args []string) error {
	config := config.Settings{}
	config.Output = "github"

	for _, reviewer := range reviewers.All {
		config.Reviewers = append(config.Reviewers, reviewer.ReviewerName())
	}

	out, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".milo.yml", out, 0666)
	if err != nil {
		return err
	}

	return nil
}
