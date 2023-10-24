package convert

func ToAny[T any, D any](source []T, transform func(T) D) []D {
	l := make([]D, 0, len(source))
	for _, v := range source {
		l = append(l, transform(v))
	}
	return l
}
