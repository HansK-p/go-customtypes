package customtypes

import (
	"testing"

	"gopkg.in/yaml.v2"
)

type SStrings struct {
	Txt0 SmartString
	Txt1 SmartString
	Tpl0 SmartString
}

func TestSmartStrings(t *testing.T) {
	yamlSmartStrings := `---
txt0: hello world
txt1: txt:hello worlds
tpl0: tpl:hello {{.Color}} world
`
	t.Log("Parsing the yaml file")
	sStrings := SStrings{}
	if err := yaml.Unmarshal([]byte(yamlSmartStrings), &sStrings); err != nil {
		t.Fatalf("Unable to unmarshal yaml data into smart strings: %s", err)
	}
	txt0 := "hello world"
	txt1 := "hello worlds"
	tpl0 := "hello blue world"
	if txt, err := sStrings.Txt0.String(); err != nil {
		t.Errorf("Got error when getting string value of txt0: %s", err)
	} else if txt != txt0 {
		t.Errorf("Result from parsing Json value txt0 is '%s', which differs from the expected value '%s'", txt, txt0)
	}
	if txt, err := sStrings.Txt1.String(); err != nil {
		t.Errorf("Got error when getting string value of txt1: %s", err)
	} else if txt != txt1 {
		t.Errorf("Result from parsing Json value txt1 is '%s', which differs from the expected value '%s'", txt, txt1)
	}
	if txt, err := sStrings.Tpl0.String(struct{ Color string }{Color: "blue"}); err != nil {
		t.Errorf("Got error when getting string value of tpl0: %s", err)
	} else if txt != tpl0 {
		t.Errorf("Result from parsing Json value tpl0 is '%s', which differs from the expected value '%s'", txt, tpl0)
	}
}
