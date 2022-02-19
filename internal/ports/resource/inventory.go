package ports

import "github.com/yemmyharry/gas-inventory-service/internal/core/domain/resource"

type Inventory interface {
	CreateItem(i *resource.Inventory) error
	UpdateItem(reference string, i *resource.Inventory) error
	DeleteItem(reference string) error
	DeleteAllUserItems(userReference string) error
	CheckItemAvailability(reference string, available string) error
}
