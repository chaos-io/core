package valid

import (
	inspection2 "github.com/chaos-io/core/valid/v2/inspection"
	rule2 "github.com/chaos-io/core/valid/v2/rule"
)

// ValueRule is an association between target value and validation rules
type ValueRule struct {
	value *inspection2.Inspected
	rules []rule2.Rule
}

// Value returns new ValueRule that passes single value through given rules
func Value(target interface{}, rules ...rule2.Rule) ValueRule {
	return ValueRule{
		value: inspection2.Inspect(target),
		rules: rules,
	}
}

// Validate runs all rules against stored value.
// It is always returns Errors error
func (v ValueRule) Validate() error {
	var errs rule2.Errors

	if v.value.Validate != nil {
		err := v.value.Validate()
		errs = append(errs, unwrapErrors(err)...)
	}

	for _, r := range v.rules {
		err := r(v.value)
		errs = append(errs, unwrapErrors(err)...)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

// unwrapErrors flattens multidimensional errors slice
func unwrapErrors(err error) []error {
	if err == nil {
		return nil
	}

	var res []error
	if multierr, ok := err.(interface{ Errors() []error }); ok {
		for _, err := range multierr.Errors() {
			res = append(res, unwrapErrors(err)...)
		}
	} else {
		res = append(res, err)
	}

	return res
}
