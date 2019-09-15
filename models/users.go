package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionUser = "Users"
)

type User struct {
	TenantID  string `json:"TenantID" bson:"TenantID" structs:"TenantID" validate:"required"`
	UserID    string `json:"UserID" bson:"UserID" structs:"UserID" validate:"required"`
	Email     string `json:"Email" bson:"Email" structs:"Email" validate:"required,email"`
	FirstName string `json:"FirstName" bson:"FirstName" structs:"FirstName" validate:"required"`
	LastName  string `json:"LastName" bson:"LastName" structs:"LastName" validate:"required"`
	TimeZone  string `json:"TimeZone" bson:"TimeZone" structs:"TimeZone,omitempty"`
	Country   string `json:"Country" bson:"Country" structs:"Country,omitempty"`
	AvatarURL string `json:"AvatarURL" bson:"AvatarURL" structs:"AvatarURL,omitempty"`
}
