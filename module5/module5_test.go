package module5

import (
	"os"
	"path"
	"testing"
)

func TestModule5InstallationOfgoimports(t *testing.T) {
	binPath := path.Join(os.Getenv("GOPATH"), "bin")
	found := FindFileAtPath(binPath, "goimports")

	if !found {
		t.Errorf("goimports cannot be found")
	}
}

func TestModule5goimportsContent(t *testing.T) {
	numOfLines := OpenFileAndCountLines("./module5_code.go")
	if numOfLines == 11 {
		t.Errorf("it looks your 'go fmt' does not work as we expected")
	}

	// TODO: add more tests
}

// FindFileAtPath returns if theFilename is found at thePath
func FindFileAtPath(thePath string, theFilename string) bool {
	if _, err := os.Stat(path.Join(thePath, theFilename)); os.IsNotExist(err) {
		return false
	}
	return true
}
