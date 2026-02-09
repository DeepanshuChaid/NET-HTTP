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

func (s *Sqlite) Create(title string, description string, completed bool) (types.Todo, error) {
	s.Db.Pr
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
