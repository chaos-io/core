package rule

import (
	"github.com/chaos-io/core/valid/v2/inspection"
)

// Rule is a validation function for a single value
type Rule = func(value *inspection.Inspected) error
