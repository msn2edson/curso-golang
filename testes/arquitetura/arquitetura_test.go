package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Nao funciona na arquitetura amd64") // go test -v
	}
	// ...
	t.Fail()
}
