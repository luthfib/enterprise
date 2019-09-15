package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionScheduledRepeating = "scheduledrepeatings"
)

// TODO: Fix types
type scheduledRepeating struct {
	TenantID string
	TaskID   string
	Date     string
	Interval string
	IsActive string
}
