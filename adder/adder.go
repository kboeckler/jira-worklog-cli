package adder

import (
	"errors"
	"fmt"
	"github.com/kboeckler/jira-worklog-cli/restclient"
	"regexp"
	"strings"
)

type Adder struct {
	client restclient.Restclient
}

func CreateAdder(client restclient.Restclient) *Adder {
	return &Adder{client: client}
}

func (a Adder) Execute(params string) (string, error) {
	words := strings.Split(params, " ")

	if strings.EqualFold(strings.TrimSpace(words[0]), "") {
		return "", errors.New("Issue cannot be null")
	}

	if len(words) < 2 {
		return "", errors.New("Worklog cannot be null")
	}

	issueKey := words[0]
	var comment string

	commentStartIndex := 2
	if len(words) >= 2 {
		for _, worklogCandidate := range words[commentStartIndex:] {
			matchesWorklogFormat, err := regexp.MatchString("\\s*[0-9]{1,2}\\s*[mMdDhH]\\s*", worklogCandidate)
			if err == nil && matchesWorklogFormat {
				commentStartIndex++
			}
		}
		if len(words) >= commentStartIndex+1 {
			comment = strings.Join(words[commentStartIndex:], " ")
		}
	}

	worklog := strings.Join(words[1:commentStartIndex], " ")

	reqBody := CreateWorklogFormatted{
		TimeSpentFormatted: worklog,
		Comment:            comment,
	}

	a.client.OpenRequestWithInput(fmt.Sprintf("/rest/api/2/issue/%s/worklog", issueKey), "POST", reqBody, nil)

	return fmt.Sprintf("Logged %s on %s (%s)\n", worklog, issueKey, comment), nil
}
