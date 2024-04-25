package jira

import (
	"log"

	jira "github.com/andygrunwald/go-jira"
)

func GetProjectKeyfromIssue(id IssueID) (ProjectID, error) {
	if val, ok := ProjectMap[IssueID(id)]; ok {
		return ProjectID(val), nil
	}

	options := &jira.GetQueryOptions{
		Fields: "project",
	}
	issue, _, err := Client.Issue.Get(string(id), options)
	if err != nil {
		log.Printf("Cant get issue %s\n ", id)
		log.Println("Cant get issue, here is why: ", err)
		return "", err
	}
	pid := ProjectID(issue.Fields.Project.Key)
	ProjectMap[IssueID(id)] = ProjectID(pid)
	return ProjectID(pid), nil
}
