package review_test

import (
	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/milo/review"
)

//Checking that review.Runner is a Runnable
var _ cmd.Runner = (*review.Runner)(nil)
