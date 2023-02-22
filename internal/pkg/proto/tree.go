package proto

import (
	"log"
)

type Pid interface {
	rune | int | int8 | byte | int16 | uint | float32 | float64 | uint32 | uint64
}

type TreeData interface {
	GetId() uint64
	GetParentId() uint64
}

type Call[T any] func(T, ...T) error

func ToTree[T TreeData](source []T, pid uint64, addChildren Call[T]) []T {
	l := make([]T, 0)
	for _, t := range source {
		if t.GetParentId() == pid {
			if child := ToTree(source, t.GetId(), addChildren); len(child) > 0 {
				if err := addChildren(t, child...); err != nil {
					log.Fatalf("Tree call run fail %v", err)
				}
			}
			l = append(l, t)
		}
	}
	return l
}
