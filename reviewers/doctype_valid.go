package reviewers

import (
	"io"
	"path/filepath"
	"strings"
)

const ValidDoctypes = []string{
	"<!DOCTYPE html>",
	`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`,
	`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">`,
}

type DoctypeValid struct{}

func (doc DoctypeValid) ReviewerName() string {
	return "doctype/valid"
}

func (doc DoctypeValid) Accepts(filePath string) bool {
	fileName := filepath.Base(filePath)
	isPartial := strings.HasPrefix(fileName, "_")
	return !isPartial
}

func (doc DoctypeValid) Review(path string, page io.Reader) ([]Fault, error) {
	return []Fault{}, nil
}
