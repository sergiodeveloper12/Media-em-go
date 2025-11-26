package main

import "fmt"

func main() {
	fmt.Println("Bem vindo aluno, ao meu sistema em GO!")
	var nome string
	fmt.Print("Digite o nome do aluno: ")
	fmt.Scanln(&nome)

	fmt.Printf("bem vindo %s ao meu sistema de notas !", nome)
	fmt.Print("insira o nome da matéria: ")
	var materia string
	fmt.Scanln(&materia)
	fmt.Printf("Agora insira a nota de %s: ", materia)
	var nota float64
	fmt.Scanln(&nota)
	fmt.Printf("insira a segunda nota de %s: ", materia)
	var nota2 float64
	fmt.Scanln(&nota2)
	media := (nota + nota2) / 2
	fmt.Printf("A média de %s em %s é: %.2f\n", nome, materia, media)
	if media >= 7 {
		fmt.Printf("Parabéns %s, você foi aprovado em %s!\n", nome, materia)
	} else {
		fmt.Printf("Infelizmente %s, você foi reprovado em %s. Tente novamente!\n", nome, materia)
	}
}
