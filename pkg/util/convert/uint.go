package convert

import "strconv"

func UnitToString(id uint) string {
	return strconv.Itoa(int(id))
}
