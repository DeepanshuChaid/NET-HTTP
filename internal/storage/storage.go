package storage

import "github.com/DeepanshuChaid/NET-HTTP.git/internal/types"

type Storage interface {
  Create(title string, description string, completed bool) (*types.Todo, error)
  
  Delete(id int) (*types.Todo, error)
  // Update()
  GetById(id int) (*types.Todo, error)
  // GetAll()
}

