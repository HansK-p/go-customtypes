package customtypes

import (
	"strings"

	"github.com/rwtodd/Go.Sed/sed"
)

type smartStringSed struct {
	b      []byte
	engine *sed.Engine
}

func (sss *smartStringSed) UnmarshalText(b []byte) (err error) {
	engine, err := sed.New(strings.NewReader(string(b)))
	if err != nil {
		return err
	}
	*sss = smartStringSed{
		b:      b,
		engine: engine,
	}
	return nil
}

func (sss *smartStringSed) MarshalText() (value []byte, err error) {
	return sss.b, nil
}

func (sss *smartStringSed) String(properties *SmartStringProperties) (result string, err error) {
	result, err = sss.engine.RunString(properties.String)
	result = strings.TrimSuffix(result, "\n")
	return
}
