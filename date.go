package customtypes

import (
	"fmt"
	"strings"
	"time"
)

// Date adds extra unmarshaling logic for time.Date
type Date time.Time

var dateLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04Z07:00",
	"2006-01-02T15:04:05",
	"2006-01-02",
}

// UnmarshalText unmarshals text into a Date
func (d *Date) UnmarshalText(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	var nt time.Time
	for _, dateLayout := range dateLayouts {
		if nt, err = time.Parse(dateLayout, s); err == nil {
			break
		}
	}
	*d = Date(nt)
	return
}

// MarshalText marshals Date as string
func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// String returns the Date in the custom format
func (d *Date) String() string {
	t := time.Time(*d)
	return fmt.Sprintf("%q", t.Format(dateLayouts[0]))
}
