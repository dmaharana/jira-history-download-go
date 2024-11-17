package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"jira-history-download/internal/config"
)

type Client struct {
	client *jira.Client
}

type HistoryItem struct {
	IssueKey    string
	Author      string
	CreatedDate string
	Field       string
	OldValue    string
	NewValue    string
}

const (
	MaxResults = 50
)

func NewClient(cfg *config.Config) (*Client, error) {
	// tp := jira.BasicAuthTransport{
	// 	Username: cfg.Username,
	// 	Password: cfg.Token,
	// }

	tp := jira.BearerAuthTransport{
		Token: cfg.Token,
	}

	client, err := jira.NewClient(tp.Client(), cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jira client: %v", err)
	}

	return &Client{client: client}, nil
}

func (c *Client) SearchIssues(jql string, startAt, maxResults int) ([]jira.Issue, error) {
	issues, _, err := c.client.Issue.Search(jql, &jira.SearchOptions{
		StartAt:    startAt,
		MaxResults: maxResults,
		Expand:     "changelog",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search issues: %v", err)
	}
	return issues, nil
}

// iterate over all issues and get their history
func (c *Client) GetHistoryForAllIssues(jql string) ([]HistoryItem, error) {
	var allHistoryItems []HistoryItem
	startAt := 0

	for {
		issues, err := c.SearchIssues(jql, startAt, MaxResults)
		if err != nil {
			return nil, fmt.Errorf("failed to search issues: %v", err)
		}

		if len(issues) == 0 {
			break
		}

		for _, issue := range issues {
			for i := 1; i < len(issue.Changelog.Histories); i++ {
				hi := HistoryItem{
					IssueKey:    issue.Key,
					Author:      issue.Changelog.Histories[i].Author.DisplayName,
					CreatedDate: issue.Changelog.Histories[i].Created,
					Field:       issue.Changelog.Histories[i].Items[0].Field,
					OldValue:    issue.Changelog.Histories[i].Items[0].FromString,
					NewValue:    issue.Changelog.Histories[i].Items[0].ToString,
				}
				allHistoryItems = append(allHistoryItems, hi)
			}
		}

		startAt += MaxResults
	}

	return allHistoryItems, nil
}

func (c *Client) GetIssueHistory(issueKey string) ([]HistoryItem, error) {
	issue, _, err := c.client.Issue.Get(issueKey, &jira.GetQueryOptions{
		Expand: "changelog",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get issue %s: %v", issueKey, err)
	}

	var historyItems []HistoryItem
	for _, history := range issue.Changelog.Histories {
		for _, item := range history.Items {
			historyItems = append(historyItems, HistoryItem{
				IssueKey:    issueKey,
				Author:      history.Author.DisplayName,
				CreatedDate: history.Created,
				Field:       item.Field,
				OldValue:    item.FromString,
				NewValue:    item.ToString,
			})
		}
	}

	return historyItems, nil
}