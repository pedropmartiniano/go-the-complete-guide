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

	printSomething(todo)
	add(1, 2)
}

// mesma coisa que any
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	var result any
	if ok {
		result = intVal + 1
	}

	floatVal, ok := value.(float64)

	if ok {
		result = floatVal + 1.5
	}

	stringVal, ok := value.(string)

	if ok {
		result = fmt.Sprintf("Valor é string: %v", stringVal)
	}
	fmt.Println(result)

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)

	case float64:
		fmt.Println("Float64: ", value)

	case string:
		fmt.Println(value)

	default:
		fmt.Println("Other value: ", value)
	}

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
