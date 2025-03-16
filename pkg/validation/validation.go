package validation

import (
	val "github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
	"github.com/umefy/go-web-app-template/pkg/validation/wrapperspb"
)

// Re-export validation functions
var (
	ValidateWithContext = val.ValidateWithContext
	ValidateByRules     = val.Validate
)

// Re-export field validation helpers
var (
	Field       = val.Field
	FieldStruct = val.FieldStruct
	Key         = val.Key
	Map         = val.Map
	Required    = val.Required
	NotNil      = val.NotNil
	Nil         = val.Nil
	Empty       = val.Empty
	Skip        = val.Skip
	By          = val.By
	WithContext = val.WithContext
	When        = val.When
	Each        = val.Each
)

// Re-export common validation rules
var (
	Length      = val.Length
	RuneLength  = val.RuneLength
	Min         = val.Min
	Max         = val.Max
	Match       = val.Match
	In          = val.In
	NotIn       = val.NotIn
	StringIn    = val.StringIn
	StringNotIn = val.StringNotIn
	Date        = val.Date
	MultipleOf  = val.MultipleOf
)

// Re-export common string format validators
var (
	IsEmail        = is.Email
	IsURL          = is.URL
	IsIP           = is.IP
	IsIPv4         = is.IPv4
	IsIPv6         = is.IPv6
	IsJSON         = is.JSON
	IsUUID         = is.UUID
	IsAlpha        = is.Alpha
	IsDigit        = is.Digit
	IsAlphanumeric = is.Alphanumeric
)

// google well-known validation rules
var (
	MinWrapperspb = wrapperspb.MinWrapperspb
	MaxWrapperspb = wrapperspb.MaxWrapperspb
)
