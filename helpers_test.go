package assert

import (
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	f := TempFile(t)
	t.Log(f)
}

func TestChdir(t *testing.T) {
	old, _ := os.Getwd()
	Chdir(t, "/")()
	//	defer Chdir(t, "/")()
	cur, _ := os.Getwd()

	Equal(t, old, cur)
}
