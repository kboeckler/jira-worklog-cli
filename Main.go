package main

import (
	"fmt"
	"github.com/jcelliott/lumber"
	"github.com/kboeckler/jira-worklog-cli/adder"
	"github.com/kboeckler/jira-worklog-cli/command"
	"github.com/kboeckler/jira-worklog-cli/lister"
	"github.com/kboeckler/jira-worklog-cli/restclient"
	"os"
	"strings"
)

func main() {
	lumber.Level(lumber.INFO)
	var client = restclient.CreateRestclient("http://localhost:8080", "test", "test")

	cmd := getArgument(1)

	var execCommand command.Command

	switch cmd {
	case "list":
		execCommand = lister.CreateLister(client)
	case "add":
		execCommand = adder.CreateAdder(client)
	default:
		fmt.Fprintf(os.Stderr, "Invalid command: %s/n", cmd)
		os.Exit(-1)
	}

	result, err := execCommand.Execute(strings.Join(os.Args[2:], " "))
	if err == nil {
		fmt.Println(result)
		os.Exit(0)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-2)
	}

}

func getArgument(position int) string {
	if len(os.Args) > position {
		return os.Args[position]
	}
	return ""
}
