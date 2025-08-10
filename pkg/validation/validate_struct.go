package validation

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	val "github.com/jellydator/validation"
)

func ValidateStruct(v interface{}, fields ...*val.FieldRules) error {
	err := val.ValidateStruct(v, fields...)

	var errs val.Errors
	if errors.As(err, &errs) {
		return &ValidateStructError{Errors: errs}
	}

	return err
}

type ValidateStructError struct {
	Errors val.Errors
}

func (e ValidateStructError) Error() string {
	es := e.Errors

	if len(es) == 0 {
		return ""
	}

	keys := make([]string, len(es))
	i := 0
	for key := range es {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	var s strings.Builder

	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}
		var validationErr *ValidateStructError
		if errors.As(es[key], &validationErr) {
			s.WriteString(fmt.Sprintf("%v.%v", key, validationErr))
		} else {
			s.WriteString(fmt.Sprintf("%v %v.", key, es[key].Error()))
		}
	}

	return s.String()
}
