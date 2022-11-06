package initialize_test

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/milo/initialize"
	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/reviewers"

	"gopkg.in/yaml.v3"
)

// Checking that initialize.Runner is a cmd.Runnable
var _ cmd.Runner = (*initialize.Runner)(nil)

func Test_InitRun(t *testing.T) {
	r := is.New(t)

	dir := os.TempDir()
	r.NoErr(os.Chdir(dir))

	c := initialize.Runner{}
	err := c.Run([]string{})
	r.NoErr(err)

	data, err := os.ReadFile(".milo.yml")
	r.NoErr(err)

	config := config.Settings{}
	r.NoErr(yaml.Unmarshal(data, &config))

	r.Equal(len(config.Reviewers), len(reviewers.All))
}
