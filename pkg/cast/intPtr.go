package cast

func Int32PtrToIntPtr(v *int32) *int {
	if v == nil {
		return nil
	}
	t := int(*v)
	return &t
}

func IntToInt32Ptr(v int) *int32 {
	v32 := int32(v)
	return &v32
}
