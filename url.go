package customtypes

import (
	"fmt"
	"net/url"
	"strings"
)

type Url url.URL

// UnmarshalText unmarshals text into an Url
func (u *Url) UnmarshalText(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	var nu *url.URL
	if nu, err = url.Parse(s); err != nil {
		return fmt.Errorf("when parsing the string '%s' as an URL: %w", s, err)
	}
	*u = Url(*nu)
	return
}

// MarshalText marshals Url as string
func (u *Url) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

// String returns the Url as a string
func (u *Url) String() string {
	tmpUrl := url.URL(*u)
	return tmpUrl.String()
}
