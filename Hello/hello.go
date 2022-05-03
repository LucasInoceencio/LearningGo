package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		println("Monitorando...")
	case 2:
		println("Exibindo logs...")
	case 0:
		println("Saindo do programa")
		os.Exit(0)
	default:
		println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1

	println("Olá, sr.", nome)
	println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	// O & está informando qual o endereço da variável comando
	//fmt.Scanf("%d", &comando)
	//println("O comando escolhido foi", comando)

	fmt.Scan(&comandoLido)
	println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	println("1 - Iniciar Monitoramento")
	println("2 - Exibir Logs")
	println("0 - Sair do Programa")
}
