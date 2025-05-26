package null

import (
	"github.com/guregu/null/v6"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// WrapperspbInt32
func ValueFromWrapperspbInt32(v *wrapperspb.Int32Value) null.Value[int] {
	return null.NewValue(int(v.GetValue()), v != nil)
}

func ValueToWrapperspbInt32(v null.Value[int]) *wrapperspb.Int32Value {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.Int32Value{
		Value: int32(v.ValueOrZero()),
	}
}

// WrapperspbInt64
func ValueFromWrapperspbInt64(v *wrapperspb.Int64Value) null.Value[int64] {
	return null.NewValue(v.GetValue(), v != nil)
}

func ValueToWrapperspbInt64(v null.Value[int64]) *wrapperspb.Int64Value {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.Int64Value{
		Value: v.ValueOrZero(),
	}
}

// WrapperspbFloat
func ValueFromWrapperspbFloat(v *wrapperspb.FloatValue) null.Value[float32] {
	return null.NewValue(v.GetValue(), v != nil)
}

func ValueToWrapperspbFloat(v null.Value[float32]) *wrapperspb.FloatValue {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.FloatValue{
		Value: v.ValueOrZero(),
	}
}

// WrapperspbDouble
func ValueFromWrapperspbDouble(v *wrapperspb.DoubleValue) null.Value[float64] {
	return null.NewValue(v.GetValue(), v != nil)
}

func ValueToWrapperspbDouble(v null.Value[float64]) *wrapperspb.DoubleValue {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.DoubleValue{
		Value: v.ValueOrZero(),
	}
}

// WrapperspbBool
func ValueFromWrapperspbBool(v *wrapperspb.BoolValue) null.Value[bool] {
	return null.NewValue(v.GetValue(), v != nil)
}

func ValueToWrapperspbBool(v null.Value[bool]) *wrapperspb.BoolValue {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.BoolValue{
		Value: v.ValueOrZero(),
	}
}

// WrapperspbString
func ValueFromWrapperspbString(v *wrapperspb.StringValue) null.Value[string] {
	return null.NewValue(v.GetValue(), v != nil)
}

func ValueToWrapperspbString(v null.Value[string]) *wrapperspb.StringValue {
	if !v.Valid {
		return nil
	}

	return &wrapperspb.StringValue{
		Value: v.ValueOrZero(),
	}
}
