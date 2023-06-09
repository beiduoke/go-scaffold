package proto

import (
	"log"
)

type Pid interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type TreeData interface {
	GetId() uint64
	GetParentId() uint64
}

func ToTree[T TreeData](source []T, pid uint64, addChildren func(T, ...T) error) []T {
	result := make([]T, 0)
	for _, t := range source {
		if t.GetParentId() == pid {
			children := ToTree(source, t.GetId(), addChildren)
			if len(children) > 0 {
				if err := addChildren(t, children...); err != nil {
					log.Fatalf("Tree call run fail %v", err)
				}
			}
			result = append(result, t)
		}
	}
	return result
}
