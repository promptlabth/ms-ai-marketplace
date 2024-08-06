package app

func PtrToType[T any](val *T) T {
	if val == nil {
		return *new(T)
	}
	return *val
}
