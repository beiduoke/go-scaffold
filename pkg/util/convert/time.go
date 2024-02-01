package convert

import "time"

func TimeValueToString(t *time.Time, format string) *string {
	if t != nil {
		s := t.Format(format)
		return &s
	}
	return nil
}
