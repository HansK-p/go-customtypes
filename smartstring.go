package customtypes

import (
	"bytes"
	"fmt"
)

type smartString interface {
	UnmarshalText(b []byte) (err error)
	MarshalText() ([]byte, error)
	String(objs ...interface{}) (string, error)
}

type SmartString struct {
	kind  string
	inner smartString
}

func (ss *SmartString) UnmarshalText(b []byte) (err error) {
	kind, value := "", b
	idx := bytes.Index(b, []byte(":"))
	if idx >= 0 {
		kind = string(b[:idx])
		value = b[idx+1:]
	}
	var inner smartString
	switch kind {
	case "":
		inner = &smartStringTxt{}
	case "txt":
		inner = &smartStringTxt{}
	case "tpl":
		inner = &smartStringTpl{}
	default:
		return fmt.Errorf("unknown smart text kind '%s'", kind)
	}
	if err := inner.UnmarshalText(value); err != nil {
		return err
	}
	*ss = SmartString{kind: kind, inner: inner}
	return nil
}

func (ss *SmartString) MarshalText() (value []byte, err error) {
	value, err = ss.inner.MarshalText()
	if ss.kind != "" {
		return ss.inner.MarshalText()
	}
	retVal := append([]byte(ss.kind), ':')
	value = append(retVal, value...)
	return
}

func (ss *SmartString) String(objs ...interface{}) (string, error) {
	return ss.inner.String(objs...)
}
