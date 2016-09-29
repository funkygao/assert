package assert

import (
	"io/ioutil"
	"os"
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

func Chdir(t *testing.T, dir string) (cleanup func()) {
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err = os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	return func() {
		os.Chdir(old)
	}
}
