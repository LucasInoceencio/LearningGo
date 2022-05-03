# LearningGo

Repositório criado para salvar informações relacionadas ao aprendizado de Go.

# Conteúdo de apoio

Site com exemplos de código para se aprender GO:

https://gobyexample.com/

Literatura de referência para Go:

https://www.casadocodigo.com.br/products/livro-google-go

Post de um dos criadores da Linguagem Go:

https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html

# Workspace Go

Go Workspace é como deve ser configurado o ambiente padrão que o Go espera.

Por padrão a pasta deve ficar na raiz do usuário com o nome "go".
	
Dentro dessa pasta deve ter outras três pastas sendos elas: "pkg", "src" e "bin".

A estrutura raiz ficará do seguinte formato:

```
usuario
    go
		pkg: Onde ficará os pacotes
		src: Onde ficará os códigos fontes
		bin: Onde vai ficar os binários compilados
```
É possível realizar a alteração dessa configuração, mas o Go vem previamente configurado nessa estrutura de pastas.

# Convenção

Uma caracterísca forte do Go é que o mesmo tem uma convenção rígida e que independente da opnião que cada um a mesma é seguida a risca.

Segue alguns itens da convenção de Go

- Não colocar ; ao final de cada instrução;
- As chaves devem ser iniciadas na frente da instrução;
- Não deve ser declaradas variavéis que não estão sendo utilizadas;
    - Neste cenário não é permitido nem a compilação do projeto;
    - Caso alguma funções retorne mais de um item e não queira utilizar um desses retornos, pode-se utilizar o ```_``` que é para ignorar essa variável;

# Um pouco da sintaxe de Go

## Declarando variáveis

```go

// Forma tradicional
var nome string 

// Forma realizando a atribuição e o próprio Go inferindo o tipo da variável de acordo com o valor atribuído
var nome = "Um nome qualquer"

// Forma utilizando o operador de atribuição curta, onde o Go infere o tipo da variável pelo valor atribuído
nome := "Mais um nome
```

## Switch Case

Switch case não tem o break, mas se colocar não faz diferença pois o compilador irá ignorar.

```go
opcao := 1

switch opcao {
case 1:
    fmt.Println("Opção 1 selecionada.")
case 2:
    fmt.Println("Opção 2 selecionada.")
case 0:
    fmt.Println("Opção 0 selecionada.")
    os.Exit(0)
default:
    println("Não conheço a opção selecionada.")
    os.Exit(-1)
}
```

## Endereço de uma variável

Para informar o endereço de uma variável basta colocar ```&``` e o nome da variável. Esse recurso é utilizado para se passar em uma função e etc, assim você faz o apontamento do local onde a informação deverá ser armazenada.

## Funções

Em Go as funções podem retornar mais de um valor, então é importante sempre tratar o retorno das funções.

```go
func MinhaPrimeiraFuncao(nome string) string{
    result := "Seja bem-vindo" + nome
    return result
}
```

## Operador discarde

O ```_``` é um identificador em branco, serve para informar para o Go ignorar uma determinada variável. Muito utilizado quando se chama uma função que retorna mais de um item, mas você só deseja manipular um desses itens.

## For

Em Go não temos while ou do while, a estrutura de iteração de resume no for que pode assumir várias formas. 

### For infinito 

```go
// Exemplo de um for infinito
for  {

} 
```

### For tradicional

```go
var sites [4]string
sites[0] = "https://random-status-code.herokuapp.com"
sites[1] = "https://www.alura.com.br"
sites[2] = "https://www.caelum.com.br"
sites[3] = "https://www.google.com.br"

for i := 0; i < len(sites); i++ {
	resp, _ := http.Get(sites[i])
	if resp.StatusCode == 200 {
		fmt.Println("Site:", sites[i], "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
	}
}
```

### For utilizando range

```go
var sites [4]string
sites[0] = "https://random-status-code.herokuapp.com"
sites[1] = "https://www.alura.com.br"
sites[2] = "https://www.caelum.com.br"
sites[3] = "https://www.google.com.br"

for i, site := range sites {
    fmt.Println("Index", i, "é referente ao site:", site)
}
```

## Array

Os arrays em Go possuem tamanho pré definidos e não é permitido a alteração desse tamanho.

```go
var sites [4]string
sites[0] = "https://random-status-code.herokuapp.com"
sites[1] = "https://www.alura.com.br"
sites[2] = "https://www.caelum.com.br"
sites[3] = "https://www.google.com.br"
```

## Slice

O Slice é uma estrutura que utiliza o array como base e uma das suas principais características é não ter um tamanho definido o que permite uma flexibilidade maior no momento da codificação.

```go
var sites []string {
    "https://random-status-code.herokuapp.com",
    "https://www.alura.com.br"
    "https://www.caelum.com.br"
}
```

## Função Len

A função len irá retornar a quantidade de itens que possue dentro de um array ou slice.

```go
var sites [5]string {
    "https://random-status-code.herokuapp.com",
    "https://www.alura.com.br"
    "https://www.caelum.com.br"
}
quantidadeSites := len(sites)
// Irá retornar 3, pois é a quantidade de itens que possui dentro do array.
```

## Função Cap

A função cap irá retornar a capacidade de um array ou slice. Pois essa estrutura pode ter uma capacidade de 4 itens mas só ter 3 itens preenchendo a estrutura.

```go
var sites [5]string {
    "https://random-status-code.herokuapp.com",
    "https://www.alura.com.br"
    "https://www.caelum.com.br"
}
capacidadeArray := cap(sites)
// Irá retornar 5, pois é a quantidade de itens que cabe dentro desse array.
```

## Função Append

Esas função permite adicionar um item em um slice ou array. Lembrando que no caso do slice quando é adicionado algo além da capacidade do mesmo, ele irá dobrar a sua capacidade.

```go
var sites []string {
    "https://random-status-code.herokuapp.com",
    "https://www.alura.com.br"
    "https://www.caelum.com.br"
}

sites = append(sites, "https://www.google.com.br")
// O slice passará a ter uma capacidade de 6 itens.
```