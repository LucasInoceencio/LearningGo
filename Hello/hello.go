package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const quantidadeMonitoramentos = 3
const delayEmSegundos = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			println("Saindo do programa")
			os.Exit(0)
		default:
			println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	// O & está informando qual o endereço da variável comando
	//fmt.Scanf("%d", &comando)
	//fmt.Println("O comando escolhido foi", comando)

	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{
		"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
		"https://www.google.com.br"}

	// For convencional
	// for i := 0; i < len(sites); i++ {
	// 	resp, _ := http.Get(sites[i])
	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
	// 	} else {
	// 		fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
	// 	}
	// }

	for i := 0; i < quantidadeMonitoramentos; i++ {
		for _, site := range sites {
			testarSite(site)
		}

		time.Sleep(delayEmSegundos * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testarSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
}
