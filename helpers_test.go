package assert

import (
	"testing"
)

func TestTempFile(t *testing.T) {
	f := TempFile(t)
	t.Log(f)
}
