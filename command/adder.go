package command

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/color"
	"github.com/kboeckler/jira-worklog-cli/restclient"
	"strings"
)

type adder interface {
	Add(params Addparams) (string, error)
}

type Addparams struct {
	IssueKey string
	Worklog  string
	Comment  string
}

type adderimpl struct {
	client restclient.Restclient
}

func CreateAdder(client restclient.Restclient) adder {
	return &adderimpl{client: client}
}

func (a adderimpl) Add(params Addparams) (string, error) {

	if strings.EqualFold(params.Worklog, "") {
		return "", errors.New("Worklog cannot be empty")
	}
	if strings.EqualFold(params.IssueKey, "") {
		return "", errors.New("IssueKey cannot be null")
	}

	reqBody := CreateWorklogFormatted{
		TimeSpentFormatted: params.Worklog,
	}
	a.client.OpenRequestWithInput(fmt.Sprintf("/rest/api/2/issue/%s/worklog", params.IssueKey), "POST", reqBody, nil)

	rows := []string{
		"IMPM-456 Testsystem verbessern  30m",
		"\\ IMPM-999 Docker               3h",
		fmt.Sprintf("%s--                              3h 30m%s", color.Yellow, color.Reset),
		"                                [...]",
		fmt.Sprintf("%s========                        4h 30m%s", color.Green, color.Reset),
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf("Logged %s on %s Docker (%s)\n", params.Worklog, params.IssueKey, params.Comment))
	res.WriteString(fmt.Sprintf("%sTotally logged today:\n%s", color.Green, color.Reset))
	for _, row := range rows {
		res.WriteString(fmt.Sprintf("%s\n", row))
	}
	return res.String(), nil
}
