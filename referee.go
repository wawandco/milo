package milo

import (
	"bytes"
	"io"
	"io/ioutil"
	"wawandco/milo/reviewers"
)

type Referee struct {
	Reviewers []Reviewer
}

var (
	DoctypePresent = "doctypepresent"
	DoctypeValid   = "doctypevalid"
	InlineCSS      = "inlinecss"
	TitlePresent   = "titlepresent"
	StyleTag       = "styletag"
	TagLowercase   = "taglowercase"
	SrcEmpty       = "srcempty"
)

func (r *Referee) Review(path string, reader io.Reader) ([]reviewers.Fault, error) {
	faults := []reviewers.Fault{}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return faults, err
	}

	for _, reviewer := range r.Reviewers {
		if !reviewer.Accepts(path) {
			continue
		}

		reader := bytes.NewReader(content)
		reviewerFaults, err := reviewer.Review(path, reader)
		if err != nil {
			return faults, err
		}

		faults = append(faults, reviewerFaults...)
	}

	return faults, nil
}

func (r *Referee) SetEnabledReviewers(enabledReviewers []string) {
	for _, reviewer := range enabledReviewers {
		switch reviewer {
		case DoctypePresent:
			r.Reviewers = append(r.Reviewers, reviewers.DoctypePresent{})
		case DoctypeValid:
			r.Reviewers = append(r.Reviewers, reviewers.DoctypeValid{})
		case InlineCSS:
			r.Reviewers = append(r.Reviewers, reviewers.InlineCSS{})
		case TitlePresent:
			r.Reviewers = append(r.Reviewers, reviewers.TitlePresent{})
		case StyleTag:
			r.Reviewers = append(r.Reviewers, reviewers.StyleTag{})
		case TagLowercase:
			r.Reviewers = append(r.Reviewers, reviewers.TagLowercase{})
		case SrcEmpty:
			r.Reviewers = append(r.Reviewers, reviewers.SrcEmpty{})
		}
	}
}

func NewReferee() *Referee {
	return &Referee{}
}
