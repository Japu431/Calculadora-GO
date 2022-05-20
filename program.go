package main

import (
	"fmt"
	"os"
)

func main() {
	println("Seja bem vindo ao programa!")

	fmt.Print("Digite um valor: ")
	var n1 int
	fmt.Scan(&n1)

	fmt.Print("Digite o segundo valor: ")
	var n2 int
	fmt.Scan(&n2)

	fmt.Println("1- Soma")
	fmt.Println("2- Subtração")
	fmt.Println("3- Multiplicação")
	fmt.Println("4- Divisão")
	fmt.Println("0- Sair do programa...")
	var operacao int
	print("Digite a operação desejada: ")
	fmt.Scan(&operacao)
	switch operacao {
	case 1:
		var res int = n1 + n2
		println("O valor de n1 + n2 : ", res)

	case 2:
		var res int = n1 - n2
		println("O valor de n1 - n2 : ", res)
	case 3:
		var res int = n1 * n2
		println("O valor de n1 * n2 : ", res)

	case 4:
		var res int = n1 / n2
		println("O valor de n1 / n2 : ", res)

	case 0:
		println("Saindo do programa...")
		os.Exit(0)

	default:
		println("Opção Inválida")
		os.Exit(-1)
	}

}
