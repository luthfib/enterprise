package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionProject = "projects"
)

type CustomField struct {
	Name string
	Type string
}

type Project struct {
	TenantID      string   `json:"TenantID" bson:"TenantID" validate:"required"`
	CreatorUserID string   `json:"CreatorUserID,omitempty" bson:"CreatorUserID,omitempty" structs:"CreatorUserID,omitempty" validate:"required"`
	ProjectID     string   `json:"ProjectID" bson:"TenantID" validate:"required"`
	Name          string   `json:"Name" bson:"Name" validate:"required"`
	Description   string   `json:"Description" bson:"Description" structs:"Description,omitempty"`
	TaskID        []string `json:"TaskID,omitempty" bson:"TaskID,omitempty" structs:"TaskID,omitempty"`
	Statuses      []string `json:"Statuses,omitempty" bson:"Statuses,omitempty" structs:"Statuses,omitempty"`
	Roles         []string `json:"Roles,omitempty" bson:"Roles,omitempty" structs:"Roles,omitempty"`
	Priorities    []string `json:"Priorities,omitempty" bson:"Priorities,omitempty" structs:"Priorities,omitempty"`
	Labels        []string `json:"Labels,omitempty" bson:"Labels,omitempty" structs:"Labels,omitempty"`
	CustomFields  []CustomField
	//CustomFieldJSON string // Todo: Make this correct type

	// User      bson.ObjectId `json:"user"`
}
