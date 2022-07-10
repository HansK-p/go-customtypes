package customtypes

import "fmt"

type smartStringTxt struct {
	b []byte
}

func (sst *smartStringTxt) UnmarshalText(b []byte) (err error) {
	*sst = smartStringTxt{b: b}
	return nil
}

func (sst *smartStringTxt) MarshalText() (value []byte, err error) {
	return sst.b, nil
}

func (sst *smartStringTxt) String(properties *SmartStringProperties) (string, error) {
	if len(sst.b) > 0 {
		return string(sst.b), nil
	}
	if properties == nil {
		return "", fmt.Errorf("properties must be set when the initial txt string is empty")
	}
	return properties.String, nil
}
