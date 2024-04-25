package util

func ValueToPtr[T any](value T) *T {
	return &value
}

func FalsePrt() *bool {
	return ValueToPtr(false)
}

func TruePrt() *bool {
	return ValueToPtr(true)
}
