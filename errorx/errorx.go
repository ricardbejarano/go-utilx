package errorx

func WrapIgnore[T any](ret T, _ error) T {
	return ret
}

func WrapPanic[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}

	return ret
}

func WrapDefault[T any](ret T, err error, def T) T {
	if err != nil {
		return def
	}

	return ret
}
