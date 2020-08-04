package reviewers

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

type TitlePresent struct{}

func (doc TitlePresent) ReviewerName() string {
	return "title/present"
}

func (doc TitlePresent) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc TitlePresent) Review(path string, page io.Reader) ([]Fault, error) {
	result := []Fault{}

	html, err := ioutil.ReadAll(page)
	if err != nil {
		return result, err
	}

	html = []byte(strings.ToLower(string(html)))
	if !strings.Contains(string(html), "<html") {
		return result, nil
	}
	re := regexp.MustCompile(`<head>[\s\w\D]*<title[^>]*>[^<]+`)
	if re.Match(html) {
		return result, nil
	}

	result = append(result, Fault{
		Reviewer: doc.ReviewerName(),
		Line:     1,
		Path:     path,
		Rule:     Rules["0004"],
	})

	return result, nil
}
