package lister

import (
	"bytes"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/color"
	"github.com/kboeckler/jira-worklog-cli/restclient"
)

type Lister struct {
	client restclient.Restclient
}

func CreateLister(client restclient.Restclient) *Lister {
	return &Lister{client: client}
}

func (l *Lister) Execute(params string) (string, error) {
	var issues List
	l.client.OpenGETRequest("/rest/api/2/search/?jql=worklogAuthor%3DcurrentUser()%20AND%20worklogDate>%3DstartOfDay()%20AND%20worklogDate<%3DendOfDay()", &issues)

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
	for _, iss := range issues.Issues {
		row := fmt.Sprintf("%s %s", iss.Key, iss.Fields.Summary)
		rows = append(rows, row)
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf("%sLogged work today:\n%s", color.Green, color.Reset))
	for _, row := range rows {
		res.WriteString(fmt.Sprintf("%s\n", row))
	}
	return res.String(), nil
}
