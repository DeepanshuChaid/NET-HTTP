package types

type Todo struct {
  Id string `json:"id"`
  Title string `json:"title" validate:"required"`
  Description string `json:"description" validate:"required"`
  Completed bool `json:"completed" validate:"required"`
}