// Initialize package is in charge of initializing the repo by generating .milo.yml
package initialize

import (
	"os"

	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/reviewers"

	"gopkg.in/yaml.v2"
)

type Runner struct{}

func (r Runner) Name() string {
	return "init"
}

func (r Runner) HelpText() string {
	return "generates .milo.yml in the current path."
}

func (r Runner) Run(args []string) error {
	c := config.Settings{}
	c.Output = "text"

	for _, reviewer := range reviewers.All {
		c.Reviewers = append(c.Reviewers, reviewer.ReviewerName())
	}

	out, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	err = os.WriteFile(".milo.yml", out, 0600)
	if err != nil {
		return err
	}

	return nil
}
