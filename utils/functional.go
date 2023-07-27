package utils

func MapSlice[T any, U any](ts []T, mapper func(T) U) []U {
	us := make([]U, 0, len(ts))

	for _, t := range ts {
		u := mapper(t)
		us = append(us, u)
	}

	return us
}
