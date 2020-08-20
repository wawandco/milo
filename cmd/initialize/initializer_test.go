package initialize_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/initialize"
	"github.com/wawandco/milo/config"
	"github.com/wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

//Checking that initialize.Runner is a cmd.Runnable
var _ cmd.Runner = (*initialize.Runner)(nil)

func Test_InitRun(t *testing.T) {
	r := require.New(t)

	dir, err := ioutil.TempDir("", "")
	r.NoError(err)
	r.NoError(os.Chdir(dir))

	c := initialize.Runner{}
	err = c.Run([]string{})
	r.NoError(err)

	r.FileExists(".milo.yml")
	data, err := ioutil.ReadFile(".milo.yml")
	r.NoError(err)

	config := config.Settings{}
	r.NoError(yaml.Unmarshal(data, &config))

	r.Len(config.Reviewers, len(reviewers.All))
}
