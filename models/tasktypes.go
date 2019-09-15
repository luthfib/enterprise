package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionTaskType = "tasktypes"
)

// Article model
type TaskType struct {
	TenantID      string // TODO: Make this correct type
	Name          string
	TaskFieldJson string // Todo: Make this correct type
}
