package review_test

import (
	"wawandco/milo/cmd"
	"wawandco/milo/cmd/review"
)

//Checking that review.Runner is a Runnable
var _ cmd.Runnable = (*review.Runner)(nil)
