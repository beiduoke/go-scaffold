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
	rune | int | int8 | byte | int16 | uint | float32 | float64
}

func ArrayUnique[T any](array []T) []T {
	if len(array) < 1 {
		return array
	}
	arrMap, arr := make(map[any]bool, 0), make([]T, 0)
	for _, v := range array {
		_, ok := arrMap[v]
		if !ok {
			arr = append(arr, v)
			arrMap[v] = true
		}
	}
	return arr
}

func ArrayStrUnique(array []string) []string {
	if len(array) < 1 {
		return array
	}
	arrMap, arr := make(map[string]bool, 0), make([]string, 0)
	for _, v := range array {
		_, ok := arrMap[v]
		if v != "" && !ok {
			arr = append(arr, v)
			arrMap[v] = true
		}
	}
	return arr
}

func ArrayNumberUnique[T Number](array []T) []T {
	if len(array) < 1 {
		return array
	}
	arrMap, arr := make(map[T]bool, 0), make([]T, 0)
	for _, v := range array {
		_, ok := arrMap[v]
		if v > 0 && !ok {
			arr = append(arr, v)
			arrMap[v] = true
		}
	}
	return arr
}

func ArrayUintToUnique(array []uint) []uint {
	arrMap := make(map[uint]uint, 0)
	for _, v := range array {
		if v <= 0 {
			continue
		}
		arrMap[v] = v
	}
	arr := make([]uint, 0)
	for _, v := range arrMap {
		arr = append(arr, v)
	}
	return arr
}

func ArrayStringToUnique(array []string) []string {
	arrMap := make(map[string]string, 0)
	for _, v := range array {
		if v == "" {
			continue
		}
		arrMap[v] = v
	}
	arr := make([]string, 0)
	for _, v := range arrMap {
		arr = append(arr, v)
	}
	return arr
}

func ArrayStringToUint(array []string) []uint {
	arrUint := make([]uint, 0, len(array))
	for _, v := range array {
		idStr, _ := strconv.Atoi(v)
		arrUint = append(arrUint, uint(idStr))
	}
	return arrUint
}
