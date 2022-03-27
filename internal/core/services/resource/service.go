package services

import (
	"gas-inventory-service/internal/core/domain/resource"
	ports "gas-inventory-service/internal/ports/resource"
)

type service struct {
	inventoryRepository ports.InventoryRepository
}

func NewService(inventoryRepository ports.InventoryRepository) *service {
	return &service{
		inventoryRepository: inventoryRepository,
	}
}

func (service *service) CreateItem(i *resource.Inventory) error {
	return service.inventoryRepository.CreateItem(i)
}

func (service *service) UpdateItem(reference string, i *resource.Inventory) error {
	return service.inventoryRepository.UpdateItem(reference, i)
}

func (service *service) DeleteItem(reference string) error {
	return service.inventoryRepository.DeleteItem(reference)
}

func (service *service) DeleteAllUserItems(userReference string) error {
	return service.inventoryRepository.DeleteAllUserItems(userReference)
}

func (service *service) SoftDeleteItem(reference string) error {
	return service.inventoryRepository.SoftDeleteItem(reference)
}

func (service *service) RestoreItem(reference string) error {
	return service.inventoryRepository.RestoreItem(reference)
}

func (service *service) SoftDeleteAllUserItems(userReference string) error {
	return service.inventoryRepository.SoftDeleteAllUserItems(userReference)
}

func (service *service) CheckItemAvailability(reference string, available string) error {
	return service.inventoryRepository.CheckItemAvailability(reference, available)
}

func (service *service) AddDocuments(reference string, documents []resource.DocumentInformation) error {
	return service.inventoryRepository.AddDocuments(reference, documents)
}

func (service *service) ValidateItem(reference string, validation resource.ValidationInformation) error {
	return service.inventoryRepository.ValidateItem(reference, validation)
}

func (service *service) GetItemDetail(reference string) (*resource.Inventory, error) {
	return service.inventoryRepository.GetItemDetail(reference)
}

func (service *service) GetItemsByMultipleSearch(search string, page int) ([]resource.Inventory, error) {
	return service.inventoryRepository.GetItemsByMultipleSearch(search, page)
}

func (service *service) GetItemsByOrganisationSearch(organisationReference string, search string, page int) ([]resource.Inventory, error) {
	return service.inventoryRepository.GetItemsByOrganisationSearch(organisationReference, search, page)
}

func (service *service) GetItemsByCategoryAndOrganisationSearch(categoryReference string, organisationReference string, search string, page int) ([]resource.Inventory, error) {
	return service.inventoryRepository.GetItemsByCategoryAndOrganisationSearch(categoryReference, organisationReference, search, page)
}

func (service *service) GetItemsByStateSearch(stateReference string, search string, page int) ([]resource.Inventory, error) {
	return service.inventoryRepository.GetItemsByStateSearch(stateReference, search, page)
}

func (service *service) GetItemsByCategoryAndStateSearch(categoryReference string, stateReference string, search string, page int) ([]resource.Inventory, error) {
	return service.inventoryRepository.GetItemsByCategoryAndStateSearch(categoryReference, stateReference, search, page)
}
