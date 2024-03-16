package main

import "fmt"

func main() {
	age := 32 // normal variable
	fmt.Println("Age: ", age)

	// var agePointer *int
	agePointer := &age // pointer variable
	fmt.Println("Age memory address: ", agePointer)
	fmt.Println("Age value from his pointer: ", *agePointer)

	adultYears := getAdultYearsValue(age)
	fmt.Println("Valor original de age(valor não alterado): ", age)
	fmt.Println("Valor alterado na variável retornada", adultYears)

	getAdultYearsReference(agePointer)
	fmt.Println("Valor original de age(valor alterado): ", age)
}

// função passando parâmetro como valor, onde será criado uma cópia da variável original com outros endereço de memória, e a variável original não será alterada
func getAdultYearsValue(age int) int {
	return age - 18
}

// função passando parâmetro como referencia(ponteiro), onde será passado o endereço de memória da variável original como parâmetro, assim o valor da variável original será alterado se ela for modificada
func getAdultYearsReference(age *int) {
	*age = *age - 18
}
