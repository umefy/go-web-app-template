package wrapperspb

import (
	val "github.com/jellydator/validation"
)

type WrapperspbRule interface {
	Error(message string) val.Rule
	val.Rule
}
