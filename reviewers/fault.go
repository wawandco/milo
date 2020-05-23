package reviewers

type Fault struct {
	ReviewerName string
	LineNumber   int
	LineContent  string

	Rule Rule
}
