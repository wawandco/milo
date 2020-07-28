package review_test

import (
	"github.com/wawandco/milo/cmd"
	"github.com/wawandco/milo/cmd/review"
)

//Checking that review.Runner is a Runnable
var _ cmd.Runnable = (*review.Runner)(nil)
