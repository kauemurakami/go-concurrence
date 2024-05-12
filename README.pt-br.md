[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-concurrence/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-concurrence/blob/main/README.md)

## Concorrência
A concorrência é um dos pontos principais do GO, uma das maiores forças da linguagem, e é por isso que ela esta sendo tão escolhido por empresas, justamente por causa da concocrrência.  
Muito facil de se utilizar, pode melhorar muito a performance da sua aplicação.<br/><br/>

A concorrência é diferente do paralelismo, o paralelismo acontece quando você tem duas ou mais tarefas sendo executadas ao mesmo tempo, isso só é possivel caso você possua um processador com mais de um nucleo, e ai sera executadas ao mesmo tempo, uma em cada nucleo.<br/><br/>

Tarefas que executam de maneira concorrênte, não estão necessariamente sendo executadas ao mesmo tempo, elas podem estar, caso haja mais de um núcleo no processador, pois seriam divididas entre os processadores, mas não necessariamente, se tiver um cpu com um nucleo só também é possivel aplicar concorrência nele, a diferença seria que você teria duas tarefas rodando, por exemplo, e uma tarefa não esperaria a outra acabar, meio que "revezariam", executam um pouco da primeira, um pouco da segunda, não necessariamente ao mesmo tempo.  

### Rotinas em GO
Go Routines são funções ou métodos que você pode chamar a execução deles e não necessariamente esperar que eles terminem para continuar com sua aplicação.  
Vejamos um exemplo de rotina em go:  
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	writePhrase("Hello world") 
	writePhrase("Olá mundo")
  //outputs
  //Hello world
  //Hello world
  //Hello world
}

func writePhrase(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
```
Ou seja, criamos uma função escrever com loop eterno com intervalo de um segundo, logo vemos que os outputs só serão gerados para a primeira função ```writePhrase("Hello world")```, como está num ```for``` infinito a segunda ```writePhrase("Olá mundo")``` nunca será executada.  
Uma ```go routine``` é uma função/método que é invocado com o prefixo ```go```, vejamos:  
```go
....
func main() {
	go writePhrase("Hello world") 
	writePhrase("Olá mundo")
  //outputs
  //Hello world
  //Hello world
  //Olá mundo
  //Olá mundo
  //....
}

func writePhrase(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
```
O que aconteceu aqui ao adicionarmos o prefixo ```go``` na chamada da função/método, você diz pro seu código o seguinte: Execute essa função, mas não espere ela acabar de executar para seguir o programa, portanto começa a executar a primeira função ,ele executa ela, mas não esperamos que ele acabe o loop infinito e pulamos pra proximo linha, continuamos o projeto. Pode ser que um termine antes do outro e etc.<br/><br/>

E caso adicionemos o prefixo ```go``` nas duas funções?  
```go
....
func main() {
	go writePhrase("Hello world") 
	go writePhrase("Olá mundo")
  //outputs
  // nada acontece
}

func writePhrase(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
```
Aqui ele começou executar a primeira função com "Hello world", não esperou terminar, passou pra próxima com "Olá mundo", só que nessa função, estamos com a mesma coisa, não esperamos terminar e seguimos o programa, mas não temos nada após essas funções, e o programa é parado.  
