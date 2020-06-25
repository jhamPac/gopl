package issue

import "time"

// IssuesURL to the github API
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult represents the results from the github issues web interface
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue represents an issue from the github API
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User represents the github user account being queried
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
