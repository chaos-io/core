package valid_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
	"github.com/chaos-io/core/xerrors"
)

func noop(_ string) error { return nil }

func TestMergeContextsSimple(t *testing.T) {
	defaultVctx := valid2.NewValidationCtx()

	myNewVctx := valid2.NewValidationCtx()
	myNewVctx.Add("simple_validator", valid2.WrapValidator(noop))

	defaultVctx.Merge(myNewVctx)

	_, ok := defaultVctx.Get("simple_validator")
	assert.True(t, ok)
}

var (
	ErrorA = xerrors.New("error A")
	ErrorB = xerrors.New("error B")
)

func generateErrorA(_ string) error { return ErrorA }
func generateErrorB(_ string) error { return ErrorB }

func TestMergeContextsReplace(t *testing.T) {
	defaultVctx := valid2.NewValidationCtx()
	defaultVctx.Add("demo_validator", valid2.WrapValidator(generateErrorA))

	myNewVctx := valid2.NewValidationCtx()
	myNewVctx.Add("demo_validator", valid2.WrapValidator(generateErrorB))

	defaultVctx.Merge(myNewVctx)

	errFunc, ok := defaultVctx.Get("demo_validator")
	assert.True(t, ok)
	assert.Equal(t, errFunc(reflect.ValueOf(""), ""), generateErrorB(""))
	assert.NotEqual(t, errFunc(reflect.ValueOf(""), ""), generateErrorA(""))
}
