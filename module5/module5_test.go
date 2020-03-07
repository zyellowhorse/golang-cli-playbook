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

// FindFileAtPath returns if theFilename is found at thePath
func FindFileAtPath(thePath string, theFilename string) bool {
	if _, err := os.Stat(path.Join(thePath, theFilename)); os.IsNotExist(err) {
		return false
	}
	return true
}
