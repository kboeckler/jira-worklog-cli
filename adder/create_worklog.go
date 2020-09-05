package adder

type CreateWorklogFormatted struct {
	TimeSpentFormatted string `json:"timeSpent"`
	Comment            string `json:"comment"`
}

type CreateWorklogInS struct {
	TimeSpentInS int `json:"timeSpentSeconds"`
}
