package schema

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Title     string `json:"title"`
}
