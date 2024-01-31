package convert

import "strconv"

func UnitToString(id uint) string {
	return strconv.Itoa(int(id))
}

func Unit32ToString(id uint32) string {
	return strconv.FormatUint(uint64(id), 10)
}

func StringToUnit32(id string) uint32 {
	ut32, _ := strconv.ParseUint(id, 10, 32)
	return uint32(ut32)
}

func Unit64ToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func StringToUnit64(id string) uint64 {
	ut32, _ := strconv.ParseUint(id, 10, 32)
	return ut32
}
