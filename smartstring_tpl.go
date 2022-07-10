package customtypes

import (
	"bytes"
	"html/template"
)

type smartStringTpl struct {
	b   []byte
	tpl *template.Template
}

func (sst *smartStringTpl) UnmarshalText(b []byte) (err error) {
	tpl, err := template.New("template").Parse(string(b))
	if err != nil {
		return err
	}
	*sst = smartStringTpl{
		b:   b,
		tpl: tpl,
	}
	return nil
}

func (sst *smartStringTpl) MarshalText() (value []byte, err error) {
	return sst.b, nil
}

func (sst *smartStringTpl) String(properties *SmartStringProperties) (result string, err error) {
	body := bytes.Buffer{}
	err = sst.tpl.Execute(&body, properties)
	return body.String(), err
}
