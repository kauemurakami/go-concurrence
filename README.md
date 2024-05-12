[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-concurrence/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-concurrence/blob/main/README.md)

## Concurrence
Competition is one of the main points of GO, one of the greatest strengths of the language, and that is why it is being chosen so much by companies, precisely because of competition.
Very easy to use, it can greatly improve the performance of your application.<br/><br/>

Concurrency is different from parallelism, parallelism happens when you have two or more tasks running at the same time, this is only possible if you have a processor with more than one core, and then they will be executed at the same time, one on each core. .<br/><br/>

Tasks that execute concurrently are not necessarily being executed at the same time, they may be, if there is more than one core in the processor, as they would be divided between the processors, but not necessarily, if you have a cpu with only one core it is also It is possible to apply concurrency to it, the difference would be that you would have two tasks running, for example, and one task would not wait for the other to finish, they would kind of "take turns", executing a little of the first, a little of the second, not necessarily at the same time.

### Routines in GO
Go Routines are functions or methods that you can call to execute them and not necessarily wait for them to finish to continue with your application.  
Let's look at an example of a go routine:  
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
In other words, we created a write function with an eternal loop with an interval of one second, so we see that the outputs will only be generated for the first function ```writePhrase("Hello world")```, as is in ```for` ``infinity the second ```writePhrase("Hello world")``` will never be executed.  
A ```go routine``` is a function/method that is invoked with the prefix ```go```, let's see:  
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
What happened here when we added the ```go``` prefix to the function/method call, you tell your code the following: Execute this function, but don't wait for it to finish executing to follow the program, so it starts executing the first function, he executes it, but we don't wait for him to finish the infinite loop and we skip to the next line, we continue the project. It could be that one finishes before the other, etc.<br/><br/>

What if we add the prefix ```go``` to both functions?  
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
Here he started executing the first function with "Hello world", he didn't wait for it to finish, he went on to the next one with "Hello world", but in this function, we have the same thing, we don't wait for it to finish and we follow the program, but we don't have anything after these functions, and the program is stopped.