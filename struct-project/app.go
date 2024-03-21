package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-project/fileops"
	"example.com/struct-project/note"
	"example.com/struct-project/todo"
)

const filePathNotes string = "./notes.json"
const filePathTodos string = "./todos.json"

func main() {
	fmt.Println("[1] - to save a To Do Item")
	fmt.Println("[2] - to save a Note")

	var choice string
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == "1" {
		saveToDo()
	} else if choice == "2" {
		saveNote()
	} else {
		panic("Invalid option")
	}

}

func saveToDo() {
	text := getUserInput("Text: ")

	todo, err := todo.New(text)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Your To Do Item has the following text:\n\n%v\n\n", todo.Text)

	todos := fileops.ReadTodoFile(filePathTodos)

	todos = append(todos, todo)

	err = fileops.WriteTodoFile(filePathTodos, todos)

	if err != nil {
		fmt.Println("There was an error when trying to write the to do item in the json file")
		return
	}

	fmt.Println("To Do saved successfully")
}

func saveNote() {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")

	note, err := note.New(title, content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
	notes := fileops.ReadNoteFile(filePathNotes)

	notes = append(notes, note)

	err = fileops.WriteNoteFile(filePathNotes, notes)

	if err != nil {
		fmt.Println("There was an error when trying to write the note in the json file")
		return
	}

	fmt.Println("Note saved successfully")
}

func getUserInput(text string) string {
	fmt.Print(text)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
