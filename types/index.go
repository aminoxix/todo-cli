package appTypes

type Todo struct {
	ID int `json:"id"`
	Task string `json:"task"`
	Checked bool `json:"checked"`
}
