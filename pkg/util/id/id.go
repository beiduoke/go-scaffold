package id

type ID[T any] interface {
	Generate() int64
	Parse() T
}
