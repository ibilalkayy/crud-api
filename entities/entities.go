package entities

type TaskVariables struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Status    string `json:"statuss"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
