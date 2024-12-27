package helper

import (
	"fmt"
	"time"
)

func FormatDate(t time.Time) string {
	return t.Format("02-01-2006")
}

func ValidateNewDate(newDate, longDate time.Time) error {
	if !newDate.After(longDate) && !newDate.Equal(longDate) {
		return fmt.Errorf("tanggal baru (%s) tidak sesuai dengan tanggal sebelumnya (%s)",
			newDate.Format("02-01-2006"), longDate.Format("02-01-2006"))
	}
	return nil
}
