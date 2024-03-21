package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface para que outros tipos/structs possam implementa-la
// toda interface que conter apenas 1 método, seu nome deve ser o nome do método + "er"
type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	todoText := getUserInput("Todo Text: ")
	title, content := getNoteData()

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething("Execução finalizada")
}

// mesma coisa que any
func printSomething(value interface{}) {
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)

	return err
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
