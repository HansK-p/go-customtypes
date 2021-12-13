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
}

// UnmarshalText unmarshals yaml into a regexp.Date
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

// MarshalText marshals regexp.Date as string
func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// String returns the time in the custom format
func (d *Date) String() string {
	t := time.Time(*d)
	return fmt.Sprintf("%q", t.Format(dateLayouts[0]))
}
