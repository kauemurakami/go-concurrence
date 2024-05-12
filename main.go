package main

import (
	"fmt"
	"time"
)

func main() {
	go writePhrase("Hello world")
	writePhrase("Olá mundo")
}

func writePhrase(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
