package models

const (
	// CollectionArticle holds the name of the articles collection
	CollectionTenant = "Tenants"
)

type Tenant struct {
	TenantID    string `json:"TenantID" bson:"TenantID" structs:"TenantID" validate:"required"`
	FirstName   string `json:"FirstName" bson:"FirstName" structs:"FirstName" validate:"required"`
	LastName    string `json:"LastName" bson:"LastName" structs:"LastName" validate:"required"`
	LicenseType string `json:"LicenseType" bson:"LicenseType" structs:"LicenseType"`
}
