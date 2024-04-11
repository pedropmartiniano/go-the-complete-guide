package main

import (
	"fmt"
	"time"
)

func greet(text string , doneChan chan bool) {
	fmt.Println(text)
	doneChan <- true
}

func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(text)
	// operador que serve para enviar um dado para um "chanel"
	doneChan <- true
	close(doneChan)
}

func main() {
	// ao adicionar a keyword "go" antes das funções, você está dizendo que quer rodar essas funções em goroutines, ou seja, irão rodar paralelamente, concorrentemente
	// ao rodar as funções assim, elas não irão bloquear as proximas execuções
	// "chan" ou "chanel" é usado para transmitir algum tipo de dado, serve para dizer que a operação está finalizada ou não
	done := make(chan bool)
	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	go greet("Hello 1", done)
	// dones[1] = make(chan bool)
	go greet("Hello 2", done)
	// dones[2] = make(chan bool)
	go slowGreet("Hello slowly", done)
	// dones[3] = make(chan bool)
	go greet("Hello 3", done)

	// essa linha diz que go só irá continuar após done receber algum tipo de dado

	// for _, done := range dones {
	// 	<- done	
	// }

	for range done {
	}
}
