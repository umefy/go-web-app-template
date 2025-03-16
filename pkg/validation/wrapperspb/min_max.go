package wrapperspb

import (
	"errors"
	"fmt"

	val "github.com/jellydator/validation"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type WrapperspbThresholdRule struct {
	threshold       interface{}
	ruleConstructor func(interface{}) val.ThresholdRule
	err             error
}

var _ val.Rule = (*WrapperspbThresholdRule)(nil)

// Validate implements validation.Rule.
func (r WrapperspbThresholdRule) Validate(value interface{}) error {
	var err error
	switch v := value.(type) {
	case *wrapperspb.Int32Value:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	case *wrapperspb.Int64Value:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	case *wrapperspb.UInt32Value:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	case *wrapperspb.UInt64Value:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	case *wrapperspb.FloatValue:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	case *wrapperspb.DoubleValue:
		err = r.ruleConstructor(r.threshold).Validate(v.GetValue())
	default:
		err = errors.New("must be a valid wrapperspb numeric value")
	}

	if err != nil {
		return r.err
	}

	return nil
}

func (r WrapperspbThresholdRule) Error(message string) val.Rule {
	r.err = errors.New(message)
	return r
}

func MinWrapperspb(threshold int) WrapperspbRule {
	return WrapperspbThresholdRule{
		threshold:       threshold,
		ruleConstructor: val.Min,
		err:             fmt.Errorf("must be greater than or equal to %d", threshold),
	}
}

func MaxWrapperspb(threshold int) WrapperspbRule {
	return WrapperspbThresholdRule{
		threshold:       threshold,
		ruleConstructor: val.Max,
		err:             fmt.Errorf("must be less than or equal to %d", threshold),
	}
}
