package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionAudit = "audits"
)

// Article model
type Audit struct {
	TenantID              string
	CollectionName        string
	ModifiedBy            string
	ModifiedFieldID       string
	ModifiedPreviousValue string
	ModifiedNewValue      string
	ModifiedDate          string
	ModifiedInstanceID    string
}
