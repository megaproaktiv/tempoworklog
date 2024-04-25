package tempo

import "time"

type Results struct {
	Attributes struct {
		Self   string `json:"self"`
		Values []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"values"`
	} `json:"attributes"`
	Author struct {
		AccountID string `json:"accountId"`
		Self      string `json:"self"`
	} `json:"author"`
	BillableSeconds int       `json:"billableSeconds"`
	CreatedAt       time.Time `json:"createdAt"`
	Description     string    `json:"description"`
	Issue           struct {
		ID   int    `json:"id"`
		Self string `json:"self"`
	} `json:"issue"`
	Self             string    `json:"self"`
	StartDate        string    `json:"startDate"`
	StartTime        string    `json:"startTime"`
	TempoWorklogID   int       `json:"tempoWorklogId"`
	TimeSpentSeconds int       `json:"timeSpentSeconds"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
type Worklog struct {
	Metadata struct {
		Count    int    `json:"count"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
	} `json:"metadata"`
	Results []Results `json:"results"`
	Self    string    `json:"self"`
}
