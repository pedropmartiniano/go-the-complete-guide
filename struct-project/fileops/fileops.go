package fileops

import (
	"encoding/json"
	"example.com/struct-project/note"
	"os"
)

func WriteJsonFile(filePath string, obj []note.Note) {
	encodedJson, _ := json.Marshal(obj)

	os.WriteFile(filePath, encodedJson, 0644)
}

func ReadJsonFile(filePath string) []note.Note {
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
