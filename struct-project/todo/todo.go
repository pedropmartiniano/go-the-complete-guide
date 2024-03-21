package todo

import (
	"errors"
	"fmt"
)

// struct com tags
type Todo struct {
	Text    string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == ""{
		return Todo{}, errors.New("it's not possible to create a todo empty")
	}

	return Todo{
		text,
	}, nil
}

func (todo Todo) Log() {
	fmt.Printf("Text: %v", todo.Text)
}
