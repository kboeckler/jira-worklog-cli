package command

import (
	"bytes"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/color"
	"github.com/kboeckler/jira-worklog-cli/issue"
)

type Lister interface {
	List() string
	SetIssues(issues issue.List)
}

type listerimpl struct {
	issues issue.List
}

func CreateLister() Lister {
	return &listerimpl{}
}

func (l *listerimpl) SetIssues(issues issue.List) {
	l.issues = issues
}

func (l *listerimpl) List() string {
	preRows := []string{
		"IMPM-123 Objekt Rückmeldung   ",
		"\\ IMPM-574 CR-Anmerkungen       1h",
		"\\ IMPM-575 Testfälle erstellen  30m",
		fmt.Sprintf("%s--                              1h 30m%s", color.Yellow, color.Reset),
		"IMPM-456 Testsystem verbessern  30m",
		"\\ IMPM-999 Docker               2h 30m",
		fmt.Sprintf("%s--                              3h%s", color.Yellow, color.Reset),
		fmt.Sprintf("%s========                        4h 30m%s", color.Green, color.Reset),
	}
	rows := make([]string, 0, len(preRows)*2)
	for _, r := range preRows {
		rows = append(rows, r)
	}
	for _, iss := range l.issues.Issues {
		row := fmt.Sprintf("%s %s", iss.Key, iss.Fields.Summary)
		rows = append(rows, row)
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf("%sLogged work today:\n%s", color.Green, color.Reset))
	for _, row := range rows {
		res.WriteString(fmt.Sprintf("%s\n", row))
	}
	return res.String()
}
