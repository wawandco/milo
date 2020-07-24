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

func (r Command) Run(args []string) error {
	c := config.Settings{}
	c.Output = "github"

	for _, reviewer := range reviewers.All {
		c.Reviewers = append(c.Reviewers, reviewer.ReviewerName())
	}

	out, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".milo.yml", out, 0600)
	if err != nil {
		return err
	}

	return nil
}
