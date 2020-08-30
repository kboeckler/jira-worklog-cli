package command

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/color"
	"strings"
)

type adder interface {
	Add(params Addparams) (string, error)
}

type Addparams struct {
	Story   string
	Worklog string
	Comment string
}

type adderimpl struct{}

func CreateAdder() adder {
	return &adderimpl{}
}

func (a adderimpl) Add(params Addparams) (string, error) {
	if strings.EqualFold(params.Worklog, "") {
		return "", errors.New("Worklog cannot be empty")
	}
	if strings.EqualFold(params.Story, "") {
		return "", errors.New("Story cannot be null")
	}

	rows := []string{
		"IMPM-456 Testsystem verbessern  30m",
		"\\ IMPM-999 Docker               3h",
		fmt.Sprintf("%s--                              3h 30m%s", color.Yellow, color.Reset),
		"                                [...]",
		fmt.Sprintf("%s========                        4h 30m%s", color.Green, color.Reset),
	}
	var res bytes.Buffer
	res.WriteString(fmt.Sprintf("Logged %s on %s Docker (%s)\n", params.Worklog, params.Story, params.Comment))
	res.WriteString(fmt.Sprintf("%sTotally logged today:\n%s", color.Green, color.Reset))
	for _, row := range rows {
		res.WriteString(fmt.Sprintf("%s\n", row))
	}
	return res.String(), nil
}
