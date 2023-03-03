package data

import "context"

type Cache[T any] interface {
	ListAllCache(context.Context) []T
	SetCache(context.Context, T) error
	GetCache(context.Context, string) T
	DeleteCache(context.Context, string) error
}
