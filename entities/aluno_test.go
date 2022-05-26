package entities

import "testing"

func TestVerificarNotas(t *testing.T) {

	aluno1 := Aluno{Nome: "Erika",
		Notas: []float64{3, 4, 2}}

	t.Run("Esperava receber score 3.0", func(t *testing.T) {
		scoreRecebido := aluno1.VerificarNotas()
		scoreEsperado := 3.0

		if scoreEsperado != scoreRecebido {
			t.Errorf(" esperado , %g, resultado  %g", scoreEsperado, scoreRecebido)
		}

	})

	aluno2 := Aluno{Nome: "Maria",
		Notas: []float64{5, 6, 5, 7}}

	t.Run("Esperava receber score 5.75", func(t *testing.T) {
		scoreRecebido := aluno2.VerificarNotas()
		scoreEsperado := 5.75

		if scoreEsperado != scoreRecebido {
			t.Errorf(" esperado , %g, resultado  %g", scoreEsperado, scoreRecebido)
		}

	})
	aluno3 := Aluno{Nome: "João",
		Notas: []float64{10, 8, 8, 10, 4}}

	t.Run("Esperava receber score 8.0", func(t *testing.T) {
		scoreRecebido := aluno3.VerificarNotas()
		scoreEsperado := 8.0

		if scoreEsperado != scoreRecebido {
			t.Errorf(" esperado  %g, resultado  %g", scoreEsperado, scoreRecebido)
		}

	})
}