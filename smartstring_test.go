package customtypes

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

type Parsed struct {
	Txt SmartString
}

type TestSet struct {
	value             string
	expect            string
	properties        *SmartStringProperties
	parseShouldError  bool
	renderShouldError bool
}

func TestSmartStrings(t *testing.T) {
	testSets := []TestSet{
		{value: `hello world`, expect: `hello world`},
		{value: `txt:hello worlds`, expect: `hello worlds`},
		{value: `tpl:hello {{.Obj.Color}} world`, expect: `hello blue world`, properties: &SmartStringProperties{Obj: struct{ Color string }{Color: "blue"}}},
		{value: `tpl:hello {{.Obj.Color}} world`, renderShouldError: true, properties: &SmartStringProperties{Obj: struct{ Form string }{Form: "rectangle"}}},
		{value: `sed:s/^(.+)\.([^.]+)$/${2}_${1}/`, expect: `hello_world`, properties: &SmartStringProperties{String: "world.hello"}},
		{value: `sed:s/^((.+)\.([^.]+)$/${2}_${1}/`, parseShouldError: true},
		{value: `sedtpl:s/^(.+)\.([^.]+)$/${2}_{{.Color}}_${1}/`, parseShouldError: true},
		{value: `awk:hello world`, parseShouldError: true},
	}
	for _, testSet := range testSets {
		yamlText := fmt.Sprintf("---\ntxt: %s", testSet.value)
		t.Logf("Parsing the yaml file based on value '%s'", testSet.value)
		parsed := Parsed{}
		if err := yaml.Unmarshal([]byte(yamlText), &parsed); err != nil {
			if !testSet.parseShouldError {
				t.Errorf("Unable to unmarshal yaml data '%s' as a smart string: %s", yamlText, err)
			}
			continue
		}
		if testSet.parseShouldError {
			t.Errorf("Parsing the yaml '%s' should have failed, but didn't", yamlText)
			continue
		}
		if txt, err := parsed.Txt.String(testSet.properties); err != nil {
			if !testSet.renderShouldError {
				t.Errorf("Got error when rendering '%#v': %s", testSet.properties, err)
			}
			continue
		} else if testSet.renderShouldError {
			t.Errorf("Rendering should have failed")
		} else if txt != testSet.expect {
			t.Errorf("Result from rendering the result set is '%s', which differs from the expected value '%s'", txt, testSet.expect)
		} else {
			t.Logf("Parsed into the expected text: %s", txt)
		}
	}
}
