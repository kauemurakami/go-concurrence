package main

import (
	"fmt"
	"time"
)

func main() {
	go writePhrase("Hello world")
	go writePhrase("Olá mundo")
}

func writePhrase(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
