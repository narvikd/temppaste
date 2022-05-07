package database

// Paste represents the table, it also includes validation directives.
type Paste struct {
	Id      string
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}
