package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaLucas := ContaCorrente{
		titular:       "Lucas",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	// contaBruna := ContaCorrente{"Bruna", 345, 654321, 200.45}

	fmt.Println(contaLucas)
	// fmt.Println(contaBruna)

	var contaCristina *ContaCorrente
	contaCristina = new(ContaCorrente)
	contaCristina.titular = "Cristina"
	contaCristina.numeroAgencia = 349
	contaCristina.numeroConta = 98765
	contaCristina.saldo = 0.80

	//var contaCristina2 *ContaCorrente
	var contaCristina2 = new(ContaCorrente)
	contaCristina2.titular = "Cristina"
	contaCristina2.numeroAgencia = 349
	contaCristina2.numeroConta = 98765
	contaCristina2.saldo = 0.80

	fmt.Println(contaCristina == contaCristina2)
}
