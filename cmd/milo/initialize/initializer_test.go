package initialize_test

import (
	"os"
	"testing"

	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/milo/initialize"
	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/reviewers"

	"gopkg.in/yaml.v3"
)

// Checking that initialize.Runner is a cmd.Runnable
var _ cmd.Runner = (*initialize.Runner)(nil)

func Test_InitRun(t *testing.T) {
	dir := os.TempDir()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("error changing directory: %v", err)
	}

	c := initialize.Runner{}
	err := c.Run([]string{})
	if err != nil {
		t.Fatalf("error running initialize: %v", err)
	}

	data, err := os.ReadFile(".milo.yml")
	if err != nil {
		t.Fatalf("error reading .milo.yml: %v", err)
	}

	config := config.Settings{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		t.Fatalf("error unmarshaling config: %v", err)
	}

	if got, want := len(config.Reviewers), len(reviewers.All); got != want {
		t.Errorf("expected %d reviewers, got %d", want, got)
	}
}
