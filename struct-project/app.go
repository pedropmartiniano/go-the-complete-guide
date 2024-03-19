package main

import (
	"example.com/struct-project/fileops"
	"example.com/struct-project/note"
	"fmt"
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

	notes := fileops.ReadJsonFile(filePath)

	notes = append(notes, note)

	fileops.WriteJsonFile(filePath, notes)
}

func getUserInput(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scanln(&input)

	return input
}
