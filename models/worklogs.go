package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionWorklog = "worklogs"
)

// TODO: Fix types
type Worklog struct {
	TenantID       string
	TaskID         string
	StartTimestamp string
	Duration       string
}
