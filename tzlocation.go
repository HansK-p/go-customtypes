package customtypes

import (
	"fmt"
	"strings"
	"time"
)

// Date adds extra unmarshaling logic for TZLocation
type TZLocation time.Location

// UnmarshalText unmarshals text into a TZLocation
func (tzl *TZLocation) UnmarshalText(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if loc, err := time.LoadLocation(s); err != nil {
		return fmt.Errorf("when converting '%s' to a location: %w", s, err)
	} else {
		*tzl = TZLocation(*loc)
	}
	return
}

// MarshalText marshals TZLocation as string
func (tzl *TZLocation) MarshalText() ([]byte, error) {
	return []byte(tzl.String()), nil
}

// String returns the TZLocation as string
func (tzl *TZLocation) String() string {
	loc := time.Location(*tzl)
	return loc.String()
}
