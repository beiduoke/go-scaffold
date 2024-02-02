package convert

import (
	"log"
	"time"
)

func TimeValueToString(t *time.Time, format string) *string {
	if t != nil {
		s := t.Format(format)
		return &s
	}
	return nil
}

func StringValueToTime(t *string, layout string) *time.Time {
	if t != nil {
		p, err := time.Parse(layout, *t)
		if err != nil {
			log.Println(err)
			return nil
		}
		return &p
	}
	return nil
}
