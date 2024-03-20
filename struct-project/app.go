package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-project/fileops"
	"example.com/struct-project/note"
)

const filePath string = "./notes.json"

func main() {
	title := getUserInput("Title: ")
	content := getUserInput("Content: ")

	note, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)

	saveNote(note)
}

func saveNote(note note.Note) {
	notes := fileops.ReadJsonFile(filePath)

	notes = append(notes, note)

	err := fileops.WriteJsonFile(filePath, notes)

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
