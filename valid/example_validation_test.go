package valid_test

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	valid2 "github.com/chaos-io/core/valid"
	xerrors2 "github.com/chaos-io/core/xerrors"
)

func Example_basicValidation() {
	// assuming you want to validate incoming user_id query arg is a valid UUID4
	var _ http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("user_id")
		if err := valid2.UUIDv4(uid); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// your business logic here

		w.WriteHeader(http.StatusOK)
	}

	// use above handler in your HTTP server instance
}

func Example_structValidation() {
	// first you need to create new validation context
	vctx := valid2.NewValidationCtx()

	// then add desired validators
	// you can use special wrapper to add predefined validators
	vctx.Add("credit_card", valid2.WrapValidator(valid2.CreditCard))

	// you can also define custom ValidatorFunc for basic types
	vctx.Add("prefixed_id", func(value reflect.Value, _ string) error {
		if value.Kind() != reflect.String {
			return valid2.ErrBadParams
		}
		sv := value.String()
		if !strings.HasPrefix(sv, "ya-") {
			return xerrors2.New("bad prefix")
		}
		return nil
	})

	// or custom types
	type YandexTeamAccount struct {
		Login    string
		LaptopID int
	}
	vctx.Add("yandex_team_account", func(value reflect.Value, _ string) error {
		acc, ok := value.Interface().(YandexTeamAccount)
		if !ok {
			return valid2.ErrBadParams
		}

		// you can use valid.Errors type for bulk errors report
		var errs valid2.Errors
		if acc.Login == "" {
			errs = append(errs, xerrors2.New("login cannot be empty"))
		}
		if acc.LaptopID < 1000 {
			errs = append(errs, xerrors2.New("laptop inventory number is too small"))
		}

		if errs != nil {
			return xerrors2.Errorf("invalid YandexTeamAccount: %w", errs)
		}
		return nil
	})

	// Validate method will be called if instance implements Validator interface
	type BadgeBalance int

	// func (bb BadgeBalance) Validate(_ *valid.ValidationCtx) (bool, error) {
	//	if bb <= 0 {
	//		return false, xerrors.New("insufficient money")
	//	}
	//	return false, nil
	// }

	// now we can validate our struct
	param := struct {
		ID           string            `valid:"prefixed_id"`
		CreditCard   string            `valid:"credit_card,omitempty"` // we can allow empty values to be provided
		StaffAccount YandexTeamAccount `valid:"yandex_team_account"`
		Badge        BadgeBalance
		SecretBadge  BadgeBalance `valid:"-"` // validation can be skipped even for Validator implementers
	}{
		ID: "ya-saint",
		StaffAccount: YandexTeamAccount{
			Login:    "saint",
			LaptopID: 42,
		},
	}

	if err := valid2.Struct(vctx, param); err != nil {
		if errs, ok := err.(valid2.Errors); ok {
			fmt.Printf("%+v", errs)
		}
	}
}
