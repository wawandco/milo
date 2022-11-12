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
)

// Checking that review.Runner is a Runnable
var _ cmd.Runner = (*review.Runner)(nil)

func TestReview(t *testing.T) {
	t.Run("run review on single file", func(t *testing.T) {
		d := t.TempDir()
		err := os.WriteFile(filepath.Join(d, "1.html"), []byte("<html><body></body></html>"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		err = r.Run([]string{
			"review",
			filepath.Join(d, "1.html"),
		})

		if err == nil {
			t.Fatal("expected an error")
		}

		if !strings.Contains(bb.String(), filepath.Join(d, "1.html")) {
			t.Errorf("Expected no faults, got %s", bb.String())
		}
	})

	t.Run("run review on folder", func(t *testing.T) {
		d := t.TempDir()

		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			err := os.WriteFile(fm, []byte("<html><body></body></html>"), 0644)
			if err != nil {
				t.Fatal(err)
			}
		}

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		err := r.Run([]string{
			"review",

			d,
		})

		if err == nil {
			t.Fatal("expected an error")
		}

		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			if !strings.Contains(bb.String(), fm) {
				t.Errorf("Expected to find %v in the output", fm)
			}
		}
	})

	t.Run("run review on multiple files", func(t *testing.T) {
		d := t.TempDir()

		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			err := os.WriteFile(fm, []byte("<html><body></body></html>"), 0644)
			if err != nil {
				t.Fatal(err)
			}
		}

		r := review.NewRunner()
		bb := bytes.NewBuffer([]byte{})
		r.SetOutput(bb)

		err := r.Run([]string{
			"review",

			filepath.Join(d, "1.html"),
			filepath.Join(d, "2.html"),
			filepath.Join(d, "3.html"),
			filepath.Join(d, "4.html"),
			filepath.Join(d, "5.html"),
		})

		if err == nil {
			t.Fatal("expected an error")
		}

		for i := 0; i < 5; i++ {
			fm := filepath.Join(d, fmt.Sprintf("%d.html", i+1))
			if !strings.Contains(bb.String(), fm) {
				t.Errorf("Expected to find %v in the output", fm)
			}
		}
	})
}
