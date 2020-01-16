package config

import (
	"testing"
)

func Test_New_Normal(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	New()
}

func Test_newConfig_Normal(t *testing.T) {
	config := newConfig("testdata/normal.yml")
	got := config.Listen.Address
	want := ":9130"

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func Test_newConfig_missingFile(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	newConfig("testdata/non_existent_file.yml")
}

func Test_newConfig_invalidYAML(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	newConfig("testdata/invalid_yaml.yml")
}
func Test_newConfig_missingPromAPIQueryURL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	newConfig("testdata/missing_prometheus_api_query_url.yml")
}

func Test_newConfig_invalidPromAPIQueryURL(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	newConfig("testdata/invalid_prometheus_api_query_url.yml")
}
