package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jcelliott/lumber"
	"github.com/kboeckler/jira-worklog-cli/command"
	"github.com/kboeckler/jira-worklog-cli/issue"
	"github.com/kboeckler/jira-worklog-cli/restclient"
	"os"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "log ", Description: "Log some time"},
		{Text: "quit", Description: "Stop application"},
		{Text: "exit", Description: "Stop Application"},
		{Text: "exit", Description: "Stop Application"},
		{Text: "stop", Description: "Stop Application"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	lumber.Level(lumber.INFO)

	cmd := getArgument(1)

	var lister = command.CreateLister()
	var adder = command.CreateAdder()
	var client = restclient.CreateRestclient("http://localhost:8080", "test", "test")
	var issues issue.List
	client.OpenGETRequest("/rest/api/2/search/?jql=worklogAuthor%3DcurrentUser()%20AND%20worklogDate>%3DstartOfDay()%20AND%20worklogDate<%3DendOfDay()", &issues)
	lister.SetIssues(issues)

	switch cmd {
	case "list":
		fmt.Println(lister.List())
		os.Exit(0)
	case "add":
		params := command.Addparams{}
		params.Story = getArgument(2)
		params.Worklog = getArgument(3)
		params.Comment = getArgument(4)
		result, err := adder.Add(params)
		if err == nil {
			fmt.Println(result)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		os.Exit(0)
	default:
		os.Exit(0)
	}

	fmt.Println("Welcome to Jira Worklog CLI")
	var loggedTime = 1
	for {
		fmt.Printf("Today's logged work time: %d\n", loggedTime)
		fmt.Println("What do you want to do?")
		t := prompt.Input("> ", completer)
		fmt.Println("You typed " + t)
		switch t {
		case "quit":
			os.Exit(0)
		case "stop":
			os.Exit(0)
		case "exit":
			os.Exit(0)
		}
		if !strings.EqualFold(strings.TrimLeft(t, "log"), t) {
			loggedTime += 1
		}
	}
}

func getArgument(position int) string {
	if len(os.Args) > position {
		return os.Args[position]
	}
	return ""
}
