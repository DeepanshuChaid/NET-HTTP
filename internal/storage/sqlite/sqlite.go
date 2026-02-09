package sqlite

import (
	"database/sql"

	"github.com/DeepanshuChaid/NET-HTTP.git/internal/config"
)

type Sqlite struct {
  Db *sql.DB
}

func New (config *config.Config) (*Sqlite, error) {
  
 _, err := sql.Open("sqlite3", config.StoragePath)
  if err != nil {
    return nil, err
  }
  
}

