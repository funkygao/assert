package assert

import (
	"io/ioutil"
	"testing"
)

// TempFile is a handy helper that returns a temporary filename.
func TempFile(t *testing.T) string {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}
