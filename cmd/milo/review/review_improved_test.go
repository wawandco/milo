package review_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/milo/review"
	"github.com/wawandco/milo/internal/assert"
)

// Checking that review.Runner is a Runnable
var _ cmd.Runner = (*review.Runner)(nil)

func TestReviewImproved(t *testing.T) {
	t.Run("run review on single file", func(t *testing.T) {
		d := t.TempDir()
		err := os.WriteFile(filepath.Join(d, "1.html"), []byte("<html><body></body></html>"), 0644)
		assert.NoError(t, err)

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		err = r.Run([]string{
			"review",
			filepath.Join(d, "1.html"),
		})

		// We expect an error because faults should be found
		assert.Error(t, err)
		assert.True(t, strings.Contains(bb.String(), filepath.Join(d, "1.html")))
	})

	t.Run("run review on folder", func(t *testing.T) {
		d := t.TempDir()

		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			err := os.WriteFile(fm, []byte("<html><body></body></html>"), 0644)
			assert.NoError(t, err)
		}

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		err := r.Run([]string{
			"review",
			d,
		})

		// We expect an error because faults should be found
		assert.Error(t, err)

		// Check that all file paths appear in the output
		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			assert.True(t, strings.Contains(bb.String(), fm), 
				"Expected to find file path %s in the output", fm)
		}
	})

	t.Run("run review on multiple files", func(t *testing.T) {
		d := t.TempDir()

		fileContents := []string{
			"<html><body></body></html>",
			"<html><body><h1>Title</h1></body></html>",
			"<html><body><p>Paragraph</p></body></html>",
			"<html><head><title>Page</title></head><body></body></html>",
			"<html><body><div></div></body></html>",
		}

		filePaths := make([]string, 5)
		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			filePaths[i] = fm
			err := os.WriteFile(fm, []byte(fileContents[i]), 0644)
			assert.NoError(t, err)
		}

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		args := append([]string{"review"}, filePaths...)
		err := r.Run(args)

		// We expect an error because faults should be found
		assert.Error(t, err)
		assert.ErrorContains(t, err, "faults found")

		// Check that all file paths appear in the output
		for _, path := range filePaths {
			assert.Contains(t, bb.String(), path)
		}
	})

	t.Run("insufficient arguments", func(t *testing.T) {
		r := review.NewRunner()
		err := r.Run([]string{"review"})
		
		assert.Error(t, err)
		assert.Equal(t, review.ErrInsufficientArgs, err)
	})

	t.Run("help text verification", func(t *testing.T) {
		r := review.NewRunner()
		assert.Equal(t, "review", r.Name())
		assert.Contains(t, r.HelpText(), "looks for faults")
		assert.NotEqual(t, "", r.HelpText())
	})

	t.Run("non-html files are skipped", func(t *testing.T) {
		d := t.TempDir()
		
		// Create an HTML file with issues
		htmlFile := filepath.Join(d, "file.html")
		err := os.WriteFile(htmlFile, []byte("<html><body></body></html>"), 0644)
		assert.NoError(t, err)
		
		// Create a non-HTML file
		txtFile := filepath.Join(d, "file.txt")
		err = os.WriteFile(txtFile, []byte("This is a text file"), 0644)
		assert.NoError(t, err)
		
		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)
		
		err = r.Run([]string{"review", d})
		
		// We should still get errors from the HTML file
		assert.Error(t, err)
		
		// The output should contain the HTML file
		assert.Contains(t, bb.String(), htmlFile)
		
		// But not the TXT file
		assert.False(t, strings.Contains(bb.String(), txtFile))
	})
}