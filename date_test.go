package customtypes

import (
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

type Dates struct {
	Rfc3339  Date
	Custom_1 Date
}

func TestDate(t *testing.T) {
	yamlDates := `---
rfc3339: "2021-12-13T23:00:05Z"
custom_1: "2021-12-13T23:00Z"
`
	t.Log("Parsing the yaml file")
	dates := Dates{}
	if err := yaml.Unmarshal([]byte(yamlDates), &dates); err != nil {
		t.Fatalf("Unable to unmarshal yaml data into dates: %s", err)
	}
	rfc3339 := time.Time(dates.Rfc3339)
	custom_1 := time.Time(dates.Custom_1)
	if rfc3339.Sub(custom_1).Seconds() != 5 {
		t.Errorf("There isn't 5 seconds difference as expected between dates, %s - %s = %f", rfc3339, custom_1, rfc3339.Sub(custom_1).Seconds())
	}
}
