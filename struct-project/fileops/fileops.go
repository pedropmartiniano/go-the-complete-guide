package fileops

import (
	"encoding/json"
	"os"

	"example.com/struct-project/note"
	"example.com/struct-project/todo"
)

func WriteNoteFile(filePath string, obj []note.Note) error {
	// essa função irá transformar o struct/obj em um json apenas dos atributos publicos deste objeto, ou seja, os atributos que está com a primeira letra em maisculo
	encodedJson, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	// Irá retornar o retorno da função os.WriteFile, que é ou um erro caso der errado, ou nil caso der certo
	return os.WriteFile(filePath, encodedJson, 0644)
}

func ReadNoteFile(filePath string) []note.Note {
	jsonBytes, err := os.Open(filePath)

	var notes []note.Note

	if err != nil {
		notes = []note.Note{}
	} else {
		decoder := json.NewDecoder(jsonBytes)
		decoder.Decode(&notes)
	}

	return notes
}

func WriteTodoFile(filePath string, obj []todo.Todo) error {
	// essa função irá transformar o struct/obj em um json apenas dos atributos publicos deste objeto, ou seja, os atributos que está com a primeira letra em maisculo
	encodedJson, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	// Irá retornar o retorno da função os.WriteFile, que é ou um erro caso der errado, ou nil caso der certo
	return os.WriteFile(filePath, encodedJson, 0644)
}

func ReadTodoFile(filePath string) []todo.Todo {
	jsonBytes, err := os.Open(filePath)

	var notes []todo.Todo

	if err != nil {
		notes = []todo.Todo{}
	} else {
		decoder := json.NewDecoder(jsonBytes)
		decoder.Decode(&notes)
	}

	return notes
}
