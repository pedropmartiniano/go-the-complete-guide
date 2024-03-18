package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// estou recebendo o ponteiro do endereço de memória da estrutura criada, em vez de uma cópia dos valores
	userPointer, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin(userFirstName, userLastName, userBirthDate, "admin@email.com", "123")

	// é possivel acessar todos os métodos de user no objeto admin, já que o user faz parte de admin
	admin.OutputUserDetails()

	// os métodos conseguem trabalhar da mesma forma com os ponteiros da estrutura, em vez do valor dela em si
	userPointer.OutputUserDetails()
	userPointer.ClearUserName()
	userPointer.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
