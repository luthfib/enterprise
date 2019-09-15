package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionTask = "tasks"
)

type Task struct {
	TenantID      string             `json:"TenantID" bson:"TenantID" structs:"TenantID"`
	TaskID        primitive.ObjectID `json:"TaskID,omitempty" bson:"TaskID" structs:"TaskID"`
	TaskType      string             `json:"TaskType,omitempty" bson:"TaskType,omitempty" structs:"TaskType,omitempty"`
	CreatorUserID string             `json:"CreatorUserID,omitempty" bson:"CreatorUserID,omitempty" structs:"CreatorUserID,omitempty"`
	Name          string             `json:"Name,omitempty" bson:"Name,omitempty" structs:"Name,omitempty"`
	Description   string             `json:"Description,omitempty" bson:"Description,omitempty" structs:"Description,omitempty"`
	CreatedTime   time.Time          `json:"CreatedTime" bson:"CreatedTime" structs:"CreatedTime"`
	Assignee      string             `json:"Assignee,omitempty" bson:"Assignee,omitempty" structs:"Assignee,omitempty"`
	TimeEstimate  string             `json:"TimeEstimate,omitempty" bson:"TimeEstimate,omitempty" structs:"TimeEstimate,omitempty"`
	IsBlockedBy   string             `json:"IsBlockedBy,omitempty" bson:"IsBlockedBy,omitempty" structs:"IsBlockedBy,omitempty"`
	Blocks        string             `json:"Blocks,omitempty" bson:"Blocks,omitempty" structs:"Blocks,omitempty"`
	StartDate     string             `json:"StartDate,omitempty" bson:"StartDate,omitempty" structs:"StartDate,omitempty"`
	EndDate       string             `json:"EndDate,omitempty" bson:"EndDate,omitempty" structs:"EndDate,omitempty"`
	Status        string             `json:"Status,omitempty" bson:"Status,omitempty" structs:"Status,omitempty"`
	Priority      string             `json:"Priority,omitempty" bson:"Priority,omitempty" structs:"Priority,omitempty"`
	IsRejected    string             `json:"IsRejected,omitempty" bson:"IsRejected,omitempty" structs:"IsRejected,omitempty"`
	Rejector      string             `json:"Rejector,omitempty" bson:"Rejector,omitempty" structs:"Rejector,omitempty"`
	ParentTaskID  string             `json:"ParentTaskID,omitempty" bson:"ParentTaskID,omitempty" structs:"ParentTaskID,omitempty"`
	Label         string             `json:"Label,omitempty" bson:"Label,omitempty" structs:"Label,omitempty"`
	Attachments   []string           `json:"Attachments,omitempty" bson:"Attachments,omitempty" structs:"Attachments,omitempty"`
	CommentIDs    []string           `json:"CommentIDs,omitempty" bson:"CommentIDs,omitempty" structs:"CommentIDs,omitempty"`
	Watchers      []string           `json:"Watchers,omitempty" bson:"Watchers,omitempty" structs:"Watchers,omitempty"`
	Checklist     []string           `json:"Checklist,omitempty" bson:"Checklist,omitempty" structs:"Checklist,omitempty"`
}
