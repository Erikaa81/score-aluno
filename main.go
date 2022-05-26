/*Criar um programa em Go que permita cadastrar um aluno curioso que quer saber sua média geral no curso (score).
Ele deve informar ao programa o nome e as notas das avaliações que ele realizou. Junto às informações do aluno, deve conter o seu score,
calculado a partir da média aritmética de quantas avaliações ele tenha feito. Crie testes para deixar ainda mais interessante o seu programa! */

package main

import (
	"fmt"

	"github.com/Erikaa81/score-aluno/entities"
)

var nome string
var nota float64

func main() {
	{
		fmt.Print("Informe seu nome: ")
		fmt.Scan(&nome)
		var notas []float64

		for {

			fmt.Print("Dígite sua nota: ")
			fmt.Scan(&nota)

			if nota == -1 {
				break
			}
			notas = append(notas, nota)

		}
		aluno := entities.Aluno{
			Nome:  nome,
			Notas: notas,
		}
		//aluno.Nome = nome
		//aluno.Notas = notas
		fmt.Printf("%s seu score é: %g\n", nome, aluno.VerificarNotas())

	}

}
