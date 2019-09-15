package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionFieldType = "fieldtypes"
)

// Article model
type FieldType struct {
	TenantID string // TODO: Make this correct type
	Name     string
	Type     string // Todo: Make this correct type
}
