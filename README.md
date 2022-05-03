# LearningGo

Repositório criado para salvar informações relacionadas ao aprendizado de Go.

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
pastaUsuario
    go
		pkg: Onde ficará os pacotes
		src: Onde ficará os códigos fontes
		bin: Onde vai ficar os binários compilados
```
É possível realizar a alteração dessa configuração, mas o Go vem previamente configurado nessa estrutura de pastas.

# Convenção

Uma caracterísca forte do Go é que o mesmo tem uma convenção rígida e que independente da opnião que cada um a mesma é seguida a risca.

Segue alguns itens da convenção de Go

- Não colocar ; ao final de cada instrução
- As chaves devem ser iniciadas na frente da instrução
- Não deve ser declaradas variavéis que não estão sendo utilizadas
    - Neste cenário não é permitido nem a compilação do projeto

# Um pouco da sintaxe de Go

## Declarando variáveis

```go
var nome string 

var nome = "Um nome qualquer"

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


:= é um operador de atribuição curto

&nome da variável está passando o endereço da variável

Em Go as funções podem retornar mais de um valor

O _ é um identificador em branco, serve para informar para o Go ignorar uma determinada variável

O for sem nenhum argumento irá rodar indefinidamente



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
quantidadeSites := len(sites)
```

## Função Cap

A função cap irá retornar a capacidade de um array ou slice. Pois essa estrutura pode ter uma capacidade de 4 itens mas só ter 3 itens preenchendo a estrutura.

```go
capacidadeArray := cap(arrayQualquer)
```