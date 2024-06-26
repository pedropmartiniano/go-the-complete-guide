package db

import (
	"database/sql" // ele pacote build-in do go utiliza o pacote go-sqlite3(ou outro) por baixo dos panos

	_ "github.com/mattn/go-sqlite3" // "_" dizer que esse pacote precisa ser usado mesmo sem interagir com ele diretamente, "_" usado para tirar o erro de importação não usada
)

var DB *sql.DB

func InitDB() {
	var err error
	// primeiro parametro é o banco de dados que será usado(mysql, postgresql...)
	// segundo parametro é a url do banco de dados que será utilizada, no caso do sqlite é somente o file path
	DB, err = sql.Open("sqlite3", "./db/api.db") // abre o db

	if err != nil {
		panic("Could not connect to database")
	}

	// controla quantas conexões pode ter aberta ao mesmo tempo desse db
	DB.SetMaxOpenConns(10)

	// quantas conexões devem ter abertas mesmo que o banco não esteja sendo usado
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTERGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table")
	}
}
