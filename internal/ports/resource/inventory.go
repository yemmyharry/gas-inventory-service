package ports

import "github.com/yemmyharry/gas-inventory-service/internal/core/domain/resource"

type Inventory interface {
	CreateItem(i *resource.Inventory) error
	UpdateItem(reference string, i *resource.Inventory) error
	DeleteItem(reference string) error
	DeleteAllUserItems(userReference string) error
	SoftDeleteItem(reference string) error
	RestoreItem(reference string) error
	SoftDeleteAllUserItems(userReference string) error
	CheckItemAvailability(reference string, available string) error
	AddDocuments(reference string, documents []resource.DocumentInformation) error
	ValidateItem(reference string, validation resource.ValidationInformation) error
	GetItemDetail(reference string) (*resource.Inventory, error)
	GetItemsByMultipleSearch(search string, page int) ([]resource.Inventory, error)
	GetItemsByOrganisationSearch(organisationReference string, search string, page int) ([]resource.Inventory, error)
	GetItemsByCategoryAndOrganisationSearch(categoryReference string, organisationReference string, search string, page int) ([]resource.Inventory, error)
	GetItemsByStateSearch(stateReference string, search string, page int) ([]resource.Inventory, error)
}
