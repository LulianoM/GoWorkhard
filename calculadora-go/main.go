package main

import (
	"fmt"
)

func main() {

	var listaMedias []int

	fmt.Println("Insira quantas notas você quer fazer a média: ")
	var numNotas int
	fmt.Scanln(&numNotas)

	i := 0
	for i < numNotas {
		i++

		var nota int
		fmt.Println("Insira nota você quer fazer a média de numero", i)
		fmt.Scanln(&nota)

		listaMedias = append(listaMedias, nota)
		fmt.Println(listaMedias)
	}

	var medias int = mean.getMean(listaMedias)
	if medias >= 7 {
		fmt.Println("Parabéns, você foi aprovade!")
	} else if medias < 5 {
		fmt.Println("Sinto muito, você foi reprovade!")
	} else {
		fmt.Println("Você está em recuperação! Estude!")
	}
}
