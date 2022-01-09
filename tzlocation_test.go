package customtypes

import (
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	tzOslo   = "Europe/Oslo"
	tzLondon = "Europe/London"
)

type Locations struct {
	Oslo   TZLocation
	London TZLocation
}

func TestLocation(t *testing.T) {
	yamlLocations := `---
oslo: "Europe/Oslo"
london: "Europe/London"
`
	t.Log("Parsing the yaml file")
	locations := Locations{}
	if err := yaml.Unmarshal([]byte(yamlLocations), &locations); err != nil {
		t.Fatalf("Unable to unmarshal yaml data into locations: %s", err)
	}
	oslo, err := time.LoadLocation(tzOslo)
	if err != nil {
		t.Fatalf("Unable to parse '%s' as a time.Location: %s", oslo, err)
	}
	if locations.Oslo.String() != oslo.String() {
		t.Errorf("Parsed Oslo location doesn't have the expected content %s != %s", locations.Oslo.String(), oslo.String())
	}

	london, err := time.LoadLocation(tzLondon)
	if err != nil {
		t.Fatalf("Unable to parse '%s' as a time.Location: %s", london, err)
	}
	if locations.London.String() != london.String() {
		t.Errorf("Parsed London location doesn't have the expected content %s != %s", locations.Oslo.String(), oslo.String())
	}
}
