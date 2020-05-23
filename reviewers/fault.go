package reviewers

type Fault struct {
	ReviewerName string
	RuleCode     string
	RuleName     string
	FileName     string
	LineNumber   int
}
