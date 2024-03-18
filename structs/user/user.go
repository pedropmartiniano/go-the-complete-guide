package user

import (
	"errors"
	"fmt"
	"time"
)

// como o tipo User está sendo exportado, ele deve estar em upperCase, porém, todos os atributos que devem ser exportados junto com o tipo, também devem estar em upperCase, caso o contrário, o atributo não será encontrado no tipo exportado
type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

// embedding User into Admin type
// é como se fosse uma classe filha da "classe" User, já que o User está sendo passado dentro dessa tipagem como um atributo, a tipagem Admin ganha acesso a todos os métodos/funções que a tipagem User possui
// é possivel realizar o embedding de forma anônima, apenas colocando o tipo do dado em vez do padrão chave e valor
// embedding anônimo: User      embedding nomeado: User User
type Admin struct {
	Email string
	Password string
	User
}

// essa função pertence somente ao tipo User, ou seja, é um método do tipo criado
// isso é possível ao definir o tipo de dado criado após a palavra reservada "func"
// portanto, como é uma função/método reservado apenas desse tipo, é possível omitir a passagem de parametro(mas somente quando a variável for definida junto com a tipagem que esse método pertence)
// receiver argument
func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

// função que recebe passagem de parametro por referencia(pointer) de uma variavel struct(objeto)
func (u *User) ClearUserName() {
	// como shortcut, em variaveis do tipo struct/objeto não é preciso converter de endereço de memória para valor, então em vez de usar (*user).firstName, é possível utilizar diretamente o user.firstName, e o go faz essa conversão implicitamente
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(firstName, lastName, birthDate, email, password string) Admin {
	user, _ := New(firstName, lastName, birthDate)
	return Admin{
		Email: email,
		Password: password,
		User: *user,
	}
}

// creation/constructor function
// uma convenção para criar estructs/objects, é ter uma função que faz a criação do objeto e o retorna
func New(userFirstName, userLastName, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("não é possível criar o objeto com parâmetros vazios")
	}

	// retornando um ponteiro do endereço de memória do objeto criado, em vez do objeto em si
	return &User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreatedAt: time.Now(),
	}, nil
	/*
		ou pode usar o shortcut que omite a chave, e instancia pela ordem dos atributos declarados no tipo, ou seja, o primeiro atributo a ser instanciado deve ser o primeiro atributo na declaração do tipo, e assim por diante

		appUser = User{
			userFirstName,
			userLastName,
			userBirthDate,
			time.Now(),
		}
	*/
}
