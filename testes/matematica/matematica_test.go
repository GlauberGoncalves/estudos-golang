package matematica

import (
	"testing"
)

const defaultError = "Valor esperado %v, mas o resultado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(defaultError, valor, valorEsperado)
	}
}
