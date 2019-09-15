package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionComment = "comments"
)

// TODO: Fix Types
type Comment struct {
	TenantID  string
	CommentID string
	UserID    string
	Text      string
	Timestamp string
}
