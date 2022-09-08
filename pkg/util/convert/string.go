package convert

import "strconv"

func StringToUint(id string) uint {
	idStr, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return uint(idStr)
}
