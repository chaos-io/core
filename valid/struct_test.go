package valid_test

import (
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"

	valid2 "github.com/chaos-io/core/valid"
	xerrors2 "github.com/chaos-io/core/xerrors"
)

var _ valid2.Validator = new(TestSelfValidator)
var _ valid2.Validator = new(TestSelfValidatorWithProceed)

type TestNullAccount struct {
	ID       int64
	Username string
	Email    string
}

func (na TestNullAccount) IsZero() bool {
	return na.ID == 0
}

type TestSelfValidator struct {
	TvmTicket []byte
}

func (sv TestSelfValidator) Validate(_ *valid2.ValidationCtx) (bool, error) {
	if len(sv.TvmTicket) < 10 &&
		sv.TvmTicket[0] != '1' {
		return false, valid2.ErrInvalidChecksum
	}
	return true, nil
}

type TestSelfValidatorWithProceed struct {
	TvmTicket []byte
	UserUID   string `valid:"uuid4"`
}

func (sv TestSelfValidatorWithProceed) Validate(_ *valid2.ValidationCtx) (bool, error) {
	if len(sv.TvmTicket) < 10 &&
		sv.TvmTicket[0] != '1' {
		return true, valid2.ErrInvalidChecksum
	}
	return true, nil
}

func validateTestNullAccount(value reflect.Value, _ string) (err error) {
	na, ok := value.Interface().(TestNullAccount)
	if !ok {
		return valid2.ErrBadParams
	}

	var errs valid2.Errors
	if err := valid2.StringLen(na.Username, 5, 255); err != nil {
		errs = append(errs, xerrors2.Errorf("Username: %w", err))
	}
	if na.Email == "" {
		errs = append(errs, xerrors2.Errorf("Email: %w", valid2.ErrEmptyString))
	}

	if len(errs) != 0 {
		err = valid2.ErrValidation.Wrap(errs)
	}

	return
}

func TestStruct_nilValidationContextPanic(t *testing.T) {
	assert.Panics(t, func() {
		_ = valid2.Struct(nil, struct {
			Name string `valid:"min=10"`
		}{
			Name: "test",
		})
	})
}

func TestStruct(t *testing.T) {
	testCases := []struct {
		name       string
		ctx        *valid2.ValidationCtx
		param      interface{}
		expectErrs valid2.Errors
	}{
		// VALID
		{
			name: "non-struct",
			ctx: func() *valid2.ValidationCtx {
				return valid2.NewValidationCtx()
			}(),
			param:      "shimba-boomba",
			expectErrs: valid2.Errors{valid2.ErrStructExpected},
		},
		{
			name: "empty_struct",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				return ctx
			}(),
			param: struct {
			}{},
			expectErrs: nil,
		},
		{
			name: "struct_with_private_fields",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				return ctx
			}(),
			param: struct {
				id string `valid:"uuid4"`
			}{
				id: uuid.Must(uuid.NewV4()).String(),
			},
			expectErrs: nil,
		},
		{
			name: "struct_with_skipped_fields",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				return ctx
			}(),
			param: struct {
				ID string `valid:"-"`
			}{
				ID: uuid.Must(uuid.NewV4()).String(),
			},
			expectErrs: nil,
		},
		{
			name: "struct_with_nil_interface",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				return ctx
			}(),
			param: struct {
				Item interface{}
			}{},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_basic_validator",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				return ctx
			}(),
			param: struct {
				ID string `valid:"uuid4"`
			}{
				ID: uuid.Must(uuid.NewV4()).String(),
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_custom_validator",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("null_account", validateTestNullAccount)
				return ctx
			}(),
			param: struct {
				ID      string          `valid:"uuid4"`
				Account TestNullAccount `valid:"null_account"`
			}{
				ID: uuid.Must(uuid.NewV4()).String(),
				Account: TestNullAccount{
					ID:       12345,
					Username: "my_long_valid_username",
					Email:    "valid_email@yandex.ru",
				},
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_empty_fields",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("credit_card", valid2.WrapValidator(valid2.CreditCard))
				ctx.Add("null_account", validateTestNullAccount)
				return ctx
			}(),
			param: struct {
				ID            string          `valid:"uuid4"`
				PaymentMethod string          `valid:"credit_card,omitempty"`
				Account       TestNullAccount `valid:"null_account,omitempty"`
				BackupCode    *string         `valid:"uuid4,omitempty"`
			}{
				ID:            uuid.Must(uuid.NewV4()).String(),
				PaymentMethod: "",
				Account:       TestNullAccount{},
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_validator_interface",
			ctx:  valid2.NewValidationCtx(),
			param: TestSelfValidator{
				TvmTicket: []byte("1234567890"),
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_param_fields",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("credit_card", valid2.WrapValidator(valid2.CreditCard))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min", valid2.Min)
				ctx.Add("max", valid2.Max)
				return ctx
			}(),
			param: struct {
				ID            string          `valid:"uuid4"`
				PaymentMethod string          `valid:"credit_card,omitempty"`
				Account       TestNullAccount `valid:"null_account,omitempty"`
				Balance       int             `valid:"min=10,max=12"`
			}{
				ID:            uuid.Must(uuid.NewV4()).String(),
				PaymentMethod: "",
				Account:       TestNullAccount{},
				Balance:       11,
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_nested_struct",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("credit_card", valid2.WrapValidator(valid2.CreditCard))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min", valid2.Min)
				ctx.Add("max", valid2.Max)
				ctx.Add("isbn", valid2.WrapValidator(valid2.ISBN))
				return ctx
			}(),
			param: struct {
				ID            string          `valid:"uuid4"`
				PaymentMethod string          `valid:"credit_card,omitempty"`
				Account       TestNullAccount `valid:"null_account,omitempty"`
				Balance       int             `valid:"min=10,max=12"`
				Metadata      struct {
					LastBook string `valid:"isbn,omitempty"`
				}
			}{
				ID:            uuid.Must(uuid.NewV4()).String(),
				PaymentMethod: "",
				Account:       TestNullAccount{},
				Balance:       11,
				Metadata: struct {
					LastBook string `valid:"isbn,omitempty"`
				}{
					LastBook: "9781234567897",
				},
			},
			expectErrs: nil,
		},
		{
			name: "valid_struct_with_iterables",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min_len", valid2.Min)
				return ctx
			}(),
			param: struct {
				DeviceUUID string              `valid:"uuid4,omitempty"`
				Accounts   []TestSelfValidator `valid:"omitempty"`
				Tags       []string            `valid:"min_len=2,omitempty"`
				Comments   map[string]string   `valid:"min_len=5,omitempty"`
			}{
				DeviceUUID: uuid.Must(uuid.NewV4()).String(),
				Accounts: []TestSelfValidator{
					{TvmTicket: []byte("1234567890")},
					{TvmTicket: []byte("1098765432")},
				},
				Tags: []string{"ololo", "trololo"},
			},
			expectErrs: nil,
		},
		// INVALID
		{
			name: "invalid_struct_with_basic_validator",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				return ctx
			}(),
			param: struct {
				ID string `valid:"uuid4"`
			}{
				ID: "some_non_uuid_string",
			},
			expectErrs: valid2.Errors{
				valid2.ErrInvalidStringLength,
			},
		},
		{
			name: "invalid_struct_with_paramed_validator",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("min", valid2.Min)
				return ctx
			}(),
			param: struct {
				ID      string
				Balance int `valid:"min=10"`
			}{
				ID:      uuid.Must(uuid.NewV4()).String(),
				Balance: 5,
			},
			expectErrs: valid2.Errors{
				valid2.ErrLesserValue,
			},
		},
		{
			name: "invalid_struct_with_paramed_validators_pair",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("min", valid2.Min)
				ctx.Add("max", valid2.Max)
				return ctx
			}(),
			param: struct {
				ID      string
				Balance int `valid:"min=10,max=15"`
				Debt    int `valid:"min=0,max=999"`
			}{
				ID:      uuid.Must(uuid.NewV4()).String(),
				Balance: 0,
				Debt:    1001,
			},
			expectErrs: valid2.Errors{
				valid2.ErrLesserValue,
				valid2.ErrGreaterValue,
			},
		},
		{
			name: "invalid_struct_with_custom_validator",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("null_account", validateTestNullAccount)
				return ctx
			}(),
			param: struct {
				ID      string          `valid:"uuid4"`
				Account TestNullAccount `valid:"null_account"`
			}{
				ID: "some_non_uuid_string",
				Account: TestNullAccount{
					ID:       12345,
					Username: "nope",
					Email:    "",
				},
			},
			expectErrs: valid2.Errors{
				valid2.ErrInvalidStringLength,
				valid2.ErrValidation.Wrap(
					valid2.Errors{
						xerrors2.Errorf("Username: %w", valid2.ErrStringTooShort),
						xerrors2.Errorf("Email: %w", valid2.ErrEmptyString),
					},
				),
			},
		},
		{
			name: "invalid_struct_with_validator_interface",
			ctx:  valid2.NewValidationCtx(),
			param: TestSelfValidator{
				TvmTicket: []byte("00000"),
			},
			expectErrs: valid2.Errors{valid2.ErrInvalidChecksum},
		},
		{
			name: "valid_struct_with_invalid_nested_struct",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("credit_card", valid2.WrapValidator(valid2.CreditCard))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min", valid2.Min)
				ctx.Add("max", valid2.Max)
				ctx.Add("isbn", valid2.WrapValidator(valid2.ISBN))
				return ctx
			}(),
			param: struct {
				ID            string          `valid:"uuid4"`
				PaymentMethod string          `valid:"credit_card,omitempty"`
				Account       TestNullAccount `valid:"null_account,omitempty"`
				Balance       int             `valid:"min=10,max=12"`
				Metadata      struct {
					LastBook string `valid:"isbn,omitempty"`
				}
			}{
				ID:            uuid.Must(uuid.NewV4()).String(),
				PaymentMethod: "",
				Account:       TestNullAccount{},
				Balance:       11,
				Metadata: struct {
					LastBook string `valid:"isbn,omitempty"`
				}{
					LastBook: "7987654312345",
				},
			},
			expectErrs: valid2.Errors{
				valid2.ErrInvalidISBN,
			},
		},
		{
			name: "invalid_struct_with_iterables",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min_len", valid2.Min)
				ctx.Add("max_len", valid2.Max)
				return ctx
			}(),
			param: struct {
				DeviceUUID string              `valid:"uuid4,omitempty"`
				Accounts   []TestSelfValidator `valid:"omitempty"`
				Tags       []string            `valid:"max_len=2,omitempty"`
				Comments   map[string]string   `valid:"min_len=2"`
			}{
				DeviceUUID: uuid.Must(uuid.NewV4()).String(),
				Accounts: []TestSelfValidator{
					{TvmTicket: []byte("trololo")},
					{TvmTicket: []byte("ololo")},
				},
				Tags: []string{"shimba", "boomba", "looken"},
			},
			expectErrs: valid2.Errors{
				valid2.ErrInvalidChecksum,
				valid2.ErrInvalidChecksum,
				valid2.ErrGreaterValue,
				valid2.ErrLesserValue,
			},
		},
		{
			name: "invalid_struct_with_proceed",
			ctx: func() *valid2.ValidationCtx {
				ctx := valid2.NewValidationCtx()
				ctx.Add("uuid4", valid2.WrapValidator(valid2.UUIDv4))
				ctx.Add("null_account", validateTestNullAccount)
				ctx.Add("min_len", valid2.Min)
				ctx.Add("max_len", valid2.Max)
				return ctx
			}(),
			param: TestSelfValidatorWithProceed{
				TvmTicket: []byte("trololo"),
				UserUID:   "ololo",
			},
			expectErrs: valid2.Errors{
				valid2.ErrInvalidChecksum,
				valid2.ErrInvalidStringLength,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errs := valid2.Struct(tc.ctx, tc.param)
			if tc.expectErrs == nil {
				assert.NoError(t, errs)
			} else {
				assert.IsType(t, valid2.Errors{}, errs)
				assert.EqualError(t, errs, tc.expectErrs.Error())
			}
		})
	}
}

func TestNestedValidationFieldErrorReturned(t *testing.T) {
	ctx := valid2.NewValidationCtx()
	ctx.Add("required", valid2.WrapValidator(func(s string) error {
		if len(s) == 0 {
			return xerrors2.New("value is required")
		}
		return nil
	}))

	type TestSelfValidatorWithTags struct {
		ID string `valid:"required"`
	}

	type TestStructWithChild struct {
		GrandChild TestSelfValidatorWithTags
	}

	type TestStructWithGrandChild struct {
		Child TestStructWithChild
	}

	param := TestStructWithGrandChild{
		Child: TestStructWithChild{
			GrandChild: TestSelfValidatorWithTags{
				ID: "",
			},
		},
	}

	errs := valid2.Struct(ctx, param).(valid2.Errors)
	if len(errs) != 1 {
		t.Error("Struct method has invalid contract")
	}

	var ferr valid2.FieldError
	if xerrors2.As(errs[0], &ferr) {
		fullPath := ferr.Path() + "." + ferr.Field()
		assert.Equal(t, fullPath, "Child.GrandChild.ID")
		assert.Equal(t, ferr.Error(), "value is required")
	} else {
		t.Error("Struct method has invalid contract")
	}
}

type SimpleTestStructWithValidate struct {
	Name string
}

func (sv SimpleTestStructWithValidate) Validate(_ *valid2.ValidationCtx) (bool, error) {
	return true, nil
}

func TestStructWithNilPointerValidation(t *testing.T) {
	type TestStructWithPointer struct {
		Child *SimpleTestStructWithValidate `valid:"omitempty"`
	}

	s := TestStructWithPointer{}

	ctx := valid2.NewValidationCtx()
	err := valid2.Struct(ctx, s)
	assert.NoError(t, err)
}
