package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("")
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

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// 	"https://www.google.com.br"}

	sites := lerSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu o erro ao fazer a requisição para o site:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo de log:", err)
	}

	fmt.Println(string(arquivo))
}

func lerSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sitesSeremMonitorados.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", err)
	}

	leitorArquivo := bufio.NewReader(arquivo)

	for {
		// As aspas simples é pra informar que é o \n em byte
		linha, err := leitorArquivo.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ocorreu um erro ao fazer a leitura do arquivo:", err)
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, online bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(online) + "\n")

	arquivo.Close()
}
