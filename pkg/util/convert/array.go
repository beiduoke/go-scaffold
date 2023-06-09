package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Map[T Number | string, D bool | byte | string | struct{}] map[T]D

func ArrayUnique[T Number | string, D bool](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	temp, result := make(Map[T, D]), make([]T, 0)
	for _, v := range arr {
		_, ok := temp[v]
		if !ok {
			result = append(result, v)
			temp[v] = true
		}
	}
	return result
}

func ArrayStringToUint(array []string) []uint {
	arrUint := make([]uint, 0, len(array))
	for _, v := range array {
		idStr, _ := strconv.Atoi(v)
		arrUint = append(arrUint, uint(idStr))
	}
	return arrUint
}

func ArrayToAny[T any, D any](source []T, transform func(T) D) []D {
	l := make([]D, 0, len(source))
	for _, v := range source {
		l = append(l, transform(v))
	}
	return l
}
