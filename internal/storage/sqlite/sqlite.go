package sqlite

import (
	"database/sql"
	"github.com/DeepanshuChaid/NET-HTTP.git/internal/types"

	"github.com/DeepanshuChaid/NET-HTTP.git/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) Create(title string, description string, completed bool) (*types.Todo, error) {

	statement, err := s.Db.Prepare("INSERT INTO todos (title, description, completed) VALUES (?, ?, ?)")
	if err != nil {
		return &types.Todo{}, err
	}

	defer statement.Close()


	result, err := statement.Exec(title, description, completed)
	if err != nil {
		return &types.Todo{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return &types.Todo{}, err
	}

	data := types.Todo{
		Id: lastId,
		Title: title,
		Description: description,
		Completed: completed,
	}
			
	return &data, nil
}

func New(config *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", config.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL UNIQUE,
		completed BOOLEAN NOT NULL
	)`)
	
	if err != nil {
		return nil, err
	}

	
	return &Sqlite{
		Db: db,
	}, nil
}

