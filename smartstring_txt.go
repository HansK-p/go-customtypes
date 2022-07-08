package customtypes

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

func (sst *smartStringTxt) String(objs ...interface{}) (string, error) {
	return string(sst.b), nil
}
