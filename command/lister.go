package command

import (
	"bytes"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/color"
)

type Lister interface {
	List() string
}

type listerimpl struct {
}

func CreateLister() Lister {
	return &listerimpl{}
}

func (*listerimpl) List() string {
	rows := []string{
		"IMPM-123   ",
		"\\ IMPM-574   1h",
		"\\ IMPM-575   30m",
		fmt.Sprintf("%s--           1h 30m%s", color.Yellow, color.Reset),
		"IMPM-456   30m",
		"\\ IMPM-999 2h 30m",
		fmt.Sprintf("%s--            3h%s", color.Yellow, color.Reset),
		fmt.Sprintf("%s========      4h 30m%s", color.Green, color.Reset),
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf("%sLogged work today:\n%s", color.Green, color.Reset))
	for _, row := range rows {
		res.WriteString(fmt.Sprintf("%s\n", row))
	}
	return res.String()
}
