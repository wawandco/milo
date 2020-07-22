package reviewers

type Fault struct {
	Reviewer string
	Path     string
	Line     int

	Rule Rule
}
