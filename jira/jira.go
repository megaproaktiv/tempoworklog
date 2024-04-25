package jira

import (
	"log"
	"math"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"github.com/joho/godotenv"
)

const jira_user_config = "JIRA_USER"
const jira_password_config = "JIRA_PASSWORD"
const jira_url_config = "JIRA_URL"

type Email string
type ProjectID string
type IssueID string
type SecondsSpent int
type AccountID string

var Client *jira.Client
var UserMap map[AccountID]Email
var ProjectMap map[IssueID]ProjectID
var UserSpentSum map[Email]SecondsSpent
var UserBillableSum map[Email]SecondsSpent

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding user's home directory:", err)
	}
	environmentFile := homeDir + string(os.PathSeparator) + ".tempoworklog"
	err = godotenv.Load(environmentFile)
	if err != nil {
		log.Fatal("Error loading .env file: ", environmentFile)
	}
	jiraURL := os.Getenv(jira_url_config)
	tp := jira.BasicAuthTransport{
		Username: os.Getenv(jira_user_config),
		Password: os.Getenv(jira_password_config),
	}
	Client, err = jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		panic(err)
	}
	// You should have only one init function in your package
	// Otherwise this would be in the user file
	UserMap = make(map[AccountID]Email)
	UserSpentSum = make(map[Email]SecondsSpent)
	UserBillableSum = make(map[Email]SecondsSpent)
	ProjectMap = make(map[IssueID]ProjectID)
}

func (s SecondsSpent) Hours() float64 {
	return math.Round(float64(s)/3600*1000) / 1000

}
