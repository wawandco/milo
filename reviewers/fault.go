package reviewers

type Fault struct {
	ReviewerName string
	Path         string
	LineNumber   int

	Rule Rule
}
