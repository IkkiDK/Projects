package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T){
	if runtime.GOARCH == "amd64"{
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}