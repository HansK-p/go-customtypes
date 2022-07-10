package customtypes

import (
	"fmt"
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

func (sss *smartStringSed) String(objs ...interface{}) (result string, err error) {
	if len(objs) != 1 {
		return "", fmt.Errorf("there must be exactly one string argument")
	}
	switch x := (objs[0]).(type) {
	case string:
		result, err = sss.engine.RunString(x)
		result = strings.TrimSuffix(result, "\n")
		return
	}
	return "", fmt.Errorf("there must be exactly one string argument - argument type received was %X", objs[0])
}
