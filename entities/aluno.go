package entities

type Aluno struct {
	Nome  string
	Notas []float64
}

func (a *Aluno) VerificarNotas() float64 {

	sum := 0.0

	for i := 0; i < len(a.Notas); i++ {

		sum += (a.Notas[i])
	}

	score := sum / float64(len(a.Notas))

	return score
}
