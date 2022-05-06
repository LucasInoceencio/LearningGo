package main

import (
	"LearningGo/Banco/clientes"
	"LearningGo/Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoBruno := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruno",
			CPF:       "123.111.123.12",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}
	fmt.Println(contaDoBruno)

	clienteAndre := clientes.Titular{Nome: "Andre", CPF: "123.123.123.12", Profissao: "Desenvolvedor Go"}
	contaDoAndre := contas.ContaPoupanca{
		Titular:       clienteAndre,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Operacao:      1}
	fmt.Println(contaDoAndre)

	contaDoAndre.Depositar(100)
	PagarBoleto(&contaDoAndre, 40)

	fmt.Println(contaDoAndre)

}
