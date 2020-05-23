package reviewers

type Fault struct {
	ReviewerName string
	LineNumber   int

	Rule Rule
}
