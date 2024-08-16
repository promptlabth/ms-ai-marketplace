package user

func TypeToPtr[T any](val T) *T {
	return &val
}
