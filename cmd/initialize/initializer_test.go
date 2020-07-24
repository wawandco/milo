package initialize_test

import (
	"io/ioutil"
	"os"
	"testing"
	"wawandco/milo/cmd"
	"wawandco/milo/cmd/initialize"
	"wawandco/milo/config"
	"wawandco/milo/reviewers"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

//Checking that init.Init is a Command
var _ cmd.Command = (*initialize.Command)(nil)

func Test_InitRun(t *testing.T) {
	r := require.New(t)

	dir, err := ioutil.TempDir("", "")
	r.NoError(err)
	r.NoError(os.Chdir(dir))

	c := initialize.Command{}
	err = c.Run([]string{})
	r.NoError(err)

	r.FileExists(".milo.yml")
	data, err := ioutil.ReadFile(".milo.yml")
	r.NoError(err)

	config := config.Settings{}
	r.NoError(yaml.Unmarshal(data, &config))

	r.Len(config.Reviewers, len(reviewers.All))
}
