package adder

import (
	"strings"
	"testing"
)

func TestExecute_EmptyCommand(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	_, err := adder.Execute("")
	if err == nil || !strings.Contains(err.Error(), "Issue") {
		t.Errorf("Expected error regarding Issue")
	}
}

func TestExecute_IssueOnlyCommand(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	_, err := adder.Execute("ABC-1")
	if err == nil || !strings.Contains(err.Error(), "Worklog") {
		t.Errorf("Expected error regarding Worklog")
	}
}

func TestExecute_SimpleCommandWithoutComment(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	adder.Execute("ABC-1 30m")
	resultInput := mock.calledInput.(CreateWorklogFormatted)
	if !strings.Contains(mock.calledUrl, "ABC-1") {
		t.Errorf("IssueKey is not ABC-1")
	}
	if resultInput.TimeSpentFormatted != "30m" {
		t.Errorf("Worklog is not 30m")
	}
}

func TestExecute_SimpleCommandWithComment(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	adder.Execute("ABC-1 30m Testcomment")
	resultInput := mock.calledInput.(CreateWorklogFormatted)
	if !strings.Contains(mock.calledUrl, "ABC-1") {
		t.Errorf("IssueKey is not ABC-1")
	}
	if resultInput.TimeSpentFormatted != "30m" {
		t.Errorf("Worklog is not 30m")
	}
	if resultInput.Comment != "Testcomment" {
		t.Errorf("Comment is not Testcomment")
	}
}

func TestExecute_CommandWithDoubleComment(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	adder.Execute("ABC-1 30m Test Comment")
	resultInput := mock.calledInput.(CreateWorklogFormatted)
	if !strings.Contains(mock.calledUrl, "ABC-1") {
		t.Errorf("IssueKey is not ABC-1")
	}
	if resultInput.TimeSpentFormatted != "30m" {
		t.Errorf("Worklog is not 30m")
	}
	if resultInput.Comment != "Test Comment" {
		t.Errorf("Comment is not Test Comment")
	}
}

func TestExecute_CommandTwoTimesWithoutComment(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	adder.Execute("ABC-1 1h 30m")
	resultInput := mock.calledInput.(CreateWorklogFormatted)
	if !strings.Contains(mock.calledUrl, "ABC-1") {
		t.Errorf("IssueKey is not ABC-1")
	}
	if resultInput.TimeSpentFormatted != "1h 30m" {
		t.Errorf("Worklog is not 1h 30m")
	}
}

func TestExecute_CommandTwoTimesWithDoubleComment(t *testing.T) {
	mock := restclientMock{}
	adder := Adder{&mock}
	adder.Execute("ABC-1 1h 30m Test Comment")
	resultInput := mock.calledInput.(CreateWorklogFormatted)
	if !strings.Contains(mock.calledUrl, "ABC-1") {
		t.Errorf("IssueKey is not ABC-1")
	}
	if resultInput.TimeSpentFormatted != "1h 30m" {
		t.Errorf("Worklog is not 1h 30m")
	}
	if resultInput.Comment != "Test Comment" {
		t.Errorf("Comment is not Test Comment")
	}
}

type restclientMock struct {
	calledUrl   string
	calledInput interface{}
}

func (r *restclientMock) OpenGETRequest(url string, record interface{}) {
	panic("implement me")
}

func (r *restclientMock) OpenDELETERequest(url string) {
	panic("implement me")
}

func (r *restclientMock) OpenRequestWithInput(url string, method string, input interface{}, response interface{}) {
	r.calledUrl = url
	r.calledInput = input
}
