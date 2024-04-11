package models

import "time"

// diz que no binding(ShouldBindJSON), somente esses atributos são obrigatórios, caso faltar algum irá ocorrer um erro
// e para os atributos que não são obrigatório, caso faltar algum, será preenchido com o "null value"
type Event struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// later add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func New() {

}
