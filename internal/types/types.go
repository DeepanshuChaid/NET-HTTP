package types

type Todo struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Completed bool `json:"completed"`
}