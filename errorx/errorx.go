package errorx

func WrapIgnore[T any](value T, _ error) T {
	return value
}

func WrapPanic[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func WrapDefault[T any](value T, err error, def T) T {
	if err != nil {
		return def
	}

	return value
}
