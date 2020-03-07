package module1

import (
	"os"
	"path"
	"testing"
)

func TestModule1CheckEnv(t *testing.T) {
	actual := os.Getenv("GOPATH")
	expected0 := ""
	expected1 := path.Join(os.Getenv("HOME"), "go")

	if actual != expected0 && actual != expected1 {
		t.Errorf("environment variable GOPATH not set properly")
	}
}
