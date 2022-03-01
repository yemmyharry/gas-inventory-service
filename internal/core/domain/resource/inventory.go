package resource

import "time"

type OrganisationInformation struct {
	OrganisationReference string `json:"organization_reference" bson:"organization_reference"`
	Industry              string `json:"industry" bson:"industry"`
	OrganisationName      string `json:"organization_name" bson:"organization_name"`
	Address               string `json:"address" bson:"address"`
	State                 string `json:"state" bson:"state"`
	Position              string `json:"position" bson:"position"`
	RCNumber              string `json:"rc_number" bson:"rc_number"`
	Location              string `json:"location" bson:"location"`
}

//type LocationInformation struct {
//	LocationReference string `json:"location_reference" bson:"location_reference"`
//	Address           string `json:"address" bson:"address"`
//	CountryReference  string `json:"country_reference" bson:"country_reference"`
//	CountryName       string `json:"country_name" bson:"country_name"`
//	StateReference    string `json:"state_reference" bson:"state_reference"`
//	StateName         string `json:"state_name" bson:"state_name"`
//	CityReference     string `json:"city_reference" bson:"city_reference"`
//	CityName          string `json:"city_name" bson:"city_name"`
//	Location          string `json:"location" bson:"location"`
//}

type CategoryInformation struct {
	CategoryReference string `json:"category_reference" bson:"category_reference"`
	CategoryName      string `json:"category_name" bson:"category_name"`
}

type SubCategoryInformation struct {
	SubCategoryReference string `json:"sub_category_reference" bson:"sub_category_reference"`
	SubCategoryName      string `json:"sub_category_name" bson:"sub_category_name"`
}

type ValidationInformation struct {
	ValidAddress  bool   `json:"valid_address" bson:"valid_address"`
	ValidPrice    bool   `json:"valid_price" bson:"valid_price"`
	ValidItem     bool   `json:"valid_item" bson:"valid_item"`
	Details       string `json:"details" bson:"details"`
	UserReference string `json:"user_reference" bson:"user_reference"`
	ValidatedBy   string `json:"validated_by" bson:"validated_by"`
}

type DocumentInformation struct {
	DocumentReference string `json:"document_reference" bson:"document_reference"`
	DocumentName      string `json:"document_name" bson:"document_name"`
}

type Inventory struct {
	Reference        string                  `json:"reference" bson:"reference"`
	UserReference    string                  `json:"user_reference" bson:"user_reference"`
	OrganizationInfo OrganisationInformation `json:"organization_info" bson:"organization_info"`
	ItemName         string                  `json:"item_name" bson:"item_name"`
	Description      string                  `json:"description" bson:"description"`
	Unit             string                  `json:"unit" bson:"unit"`
	CategoryInfo     CategoryInformation     `json:"category_info" bson:"category_info"`
	SubCategoryInfo  SubCategoryInformation  `json:"sub_category_info" bson:"sub_category_info"`
	Sponsored        bool                    `json:"sponsored" bson:"sponsored"`
	Available        bool                    `json:"available" bson:"available"`
	IsValidatedItem  bool                    `json:"is_validated_item" bson:"is_validated_item"`
	ValidationInfo   ValidationInformation   `json:"validation_info" bson:"validation_info"`
	DocumentInfo     []DocumentInformation   `json:"document_info" bson:"document_info"`
	NoOfDocuments    int                     `json:"no_of_documents" bson:"no_of_documents"`
	IsDeleted        bool                    `json:"is_deleted" bson:"is_deleted"`
	CreatedAt        time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at" bson:"updated_at"`
	DeletedAt        time.Time               `json:"deleted_at" bson:"deleted_at"`
}
