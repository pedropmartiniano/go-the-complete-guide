package main

import "fmt"

// custom types to add methods to it
type str string

func (text str) log() {
	fmt.Println(text)
}

func costumTypes(){
	var name str = "Pedro"

	name.log()
}