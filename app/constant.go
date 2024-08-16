package app

func PtrToType[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}

func TypeToPtr[T any](val T) *T {
	return &val
}
