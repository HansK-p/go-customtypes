package customtypes

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

type Urls struct {
	Urls []Url `yaml:"urls"`
}

func TestUrl(t *testing.T) {
	strUrls := []string{
		"https://www.google.com",
		"/test/file.json",
		"file://test/file.json",
		"hello://dadada-dotdot:33:33",
	}
	yamlUrls := `---
urls:`
	for _, strUrl := range strUrls {
		yamlUrls = fmt.Sprintf("%s\n- %s", yamlUrls, strUrl)
	}
	t.Logf("Parsing the yaml file:\n%s", yamlUrls)
	urls := Urls{}
	if err := yaml.Unmarshal([]byte(yamlUrls), &urls); err != nil {
		t.Fatalf("Unable to unmarshal yaml data into URLs: %s", err)
	}
	for idx, strUrl := range strUrls {
		if strUrl != urls.Urls[idx].String() {
			t.Errorf("Yaml url #%d '%s' parsed value is '%s'", idx, strUrl, urls.Urls[idx].String())
		}
	}
}
