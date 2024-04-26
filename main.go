package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"twl/dates"
	"twl/jira"
	"twl/tempo"
)

// Data model for the API response.
type ApiResponse struct {
	Message string `json:"message"`
}

func main() {
	// Defining flags.
	period := flag.String("period", "week", "Choose the period: week, month, lastmonth")
	project := flag.String("project", "", "Filter by project")
	long := flag.Bool("long", false, "Show whole worklog entries")
	// outputFormat := flag.String("output", "text", "Choose the output format: text, json, csv")
	flag.Parse()

	if *period != "week" && *period != "month" && *period != "lastmonth" {
		fmt.Println("Invalid period. Choose one of: thisweek, thismonth, lastmonth")
		return
	}
	dateRange := dates.GetDateRange(*period)
	fmt.Printf("From: %s, To: %s\n", dateRange.From, dateRange.To)

	response, err := tempo.CallTempoAPI(dateRange.From, dateRange.To)

	if err != nil {
		fmt.Println("Error calling Tempo API:", err)
		return
	}
	raw := false
	if raw {
		fmt.Println(response)
		return
	}
	var worklogs tempo.Worklog
	// Unmarshal the JSON data into a Worklog struct.
	// The Worklog struct is defined in tempo.go.
	//
	// The JSON data is an array of worklogs, but we only need the first one.
	err = json.Unmarshal([]byte(response), &worklogs)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	allResults := []tempo.Results{}
	allResults = append(allResults, worklogs.Results...)
	safetyCounter := 1000
	counter := 0
	limit := worklogs.Metadata.Limit
	increment := limit
	answerLength := limit
	for answerLength > 0 {
		counter++
		response, err = tempo.CallTempoNext(worklogs.Metadata.Next)
		if counter > safetyCounter {
			break
		}
		if answerLength < increment {
			break
		}
		if err != nil {
			fmt.Println("Error calling Tempo API:", err)
			return
		}
		err = json.Unmarshal([]byte(response), &worklogs)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		allResults = append(allResults, worklogs.Results...)
		answerLength = len(worklogs.Results)

	}
	fmt.Printf("Project Key,\t User,\tDate,\tIssue, \tTime Spent,\t Billable\tDescription\n")
	sumSpent := 0
	sumBillable := 0
	for _, work := range allResults {
		hours := secondsToHours(work.TimeSpentSeconds)
		hoursBillable := secondsToHours(work.BillableSeconds)
		accountId := work.Author.AccountID
		userName, err := jira.GetUserFromAccount(jira.AccountID(accountId))
		issueId := strconv.Itoa(work.Issue.ID)
		projectId, err := jira.GetProjectKeyfromIssue(jira.IssueID(issueId))
		startDate := work.StartDate
		description := strings.ReplaceAll(work.Description, "\n", "")
		max := 32
		if !*long {
			if len(description) > max {
				description = description[:max]
			}
		}

		if err != nil {
			fmt.Println("Error getting projectkey:", err)
			projectId = "Unknown"
		}
		if *project == "" || (string(projectId) == *project) {
			fmt.Printf("%s,\t%s,\t%s,\t%s,\t%s,\t%s,\t%s \n",
				projectId,
				userName,
				startDate,
				issueId,
				hours,
				hoursBillable,
				description)
			sumBillable += work.BillableSeconds
			sumSpent += work.BillableSeconds
		}
	}

	fmt.Printf("SumSpent: %s, SumBillable: %s\n",
		secondsToHours(sumSpent),
		secondsToHours(sumBillable))
}
func secondsToHours(seconds int) string {
	hours := float64(seconds) / 3600.0
	return fmt.Sprintf("%.2f", hours)
}
