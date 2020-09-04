package issue

type List struct {
	StartAt int     `json:"startAt"`
	Total   int     `json:"total"`
	Issues  []Issue `json:"issues"`
}

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Summary string `json:"summary"`
}
