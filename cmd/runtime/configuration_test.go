package runtime

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var configTemplate = `
output: "github"
reviewers:
  - doctype/present 
`

func Test_LoadConfiguration(t *testing.T) {
	r := require.New(t)
	wd, err := ioutil.TempDir("", "")
	r.NoError(err)

	cwd, err := os.Getwd()
	r.NoError(err)
	defer func() {
		os.Chdir(cwd)
		os.RemoveAll(wd)
	}()

	err = ioutil.WriteFile(filepath.Join(wd, ".milo.yml"), []byte(configTemplate), 0777)
	r.NoError(err)

	os.Chdir(wd)

	config := LoadConfiguration()
	r.Len(config.Reviewers, 1)
	r.Len(config.SelectedReviewers(), 1)

	err = os.Remove(filepath.Join(wd, ".milo.yml"))
	r.NoError(err)

	config = LoadConfiguration()
	r.Len(config.Reviewers, 0)
	r.Len(config.SelectedReviewers(), 9)
}
