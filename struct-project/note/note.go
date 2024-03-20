package note

import (
	"errors"
	"fmt"
	"time"
)

// struct com tags
type Note struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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
	fmt.Printf("Title: %v\nContent: %v\nCreated_at: %v", note.Title, note.Content, note.CreatedAt)
}
