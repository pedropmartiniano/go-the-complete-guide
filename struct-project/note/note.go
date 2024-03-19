package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title      string
	Content    string
	Created_at time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("it's not possible to create a note with a title or content empty")
	}

	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

func (note Note) Log() {
	fmt.Printf("Title: %v\nContent: %v\nCreated_at: %v", note.Title, note.Content, note.Created_at)
}
