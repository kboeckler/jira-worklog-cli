package command

type CreateWorklogFormatted struct {
	TimeSpentFormatted string `json:"timeSpent"`
}

type CreateWorklogInS struct {
	TimeSpentInS int `json:"timeSpentSeconds"`
}
