package null

import (
	"time"

	"github.com/guregu/null/v6"
)

type BasicType interface {
	~string | ~int | ~int16 | ~int32 | ~int64 | ~float64 | ~bool | time.Time
}

func NewValue[T BasicType](v T, valid bool) null.Value[T] {
	return null.NewValue(v, valid)
}

func ValueFrom[T BasicType](v T) null.Value[T] {
	return null.ValueFrom(v)
}

func ValueFromPtr[T BasicType](v *T) null.Value[T] {
	return null.ValueFromPtr(v)
}
