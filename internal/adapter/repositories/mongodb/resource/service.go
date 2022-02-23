package resource

import (
	"github.com/globalsign/mgo/bson"
	"github.com/yemmyharry/gas-inventory-service/internal/core/domain/resource"
	"log"
	"strconv"
	"time"
)

func (r *MongoRepository) CreateItem(i *resource.Inventory) error {
	err := r.Client.C("items").Insert(i)
	if err != nil {
		log.Println("unable to insert data:", err.Error())
	}
	return err
}

func (r *MongoRepository) UpdateItem(reference string, i *resource.Inventory) error {
	locationInformation := resource.LocationInformation{
		LocationReference: i.LocationInfo.LocationReference,
		Address:           i.LocationInfo.Address,
		CountryReference:  i.LocationInfo.CountryReference,
		CountryName:       i.LocationInfo.CountryName,
		StateReference:    i.LocationInfo.StateReference,
		StateName:         i.LocationInfo.StateName,
		CityReference:     i.LocationInfo.CityReference,
		CityName:          i.LocationInfo.CityName,
		Location:          i.LocationInfo.Location,
	}

	categoryInformation := resource.CategoryInformation{
		CategoryReference: i.CategoryInfo.CategoryReference,
		CategoryName:      i.CategoryInfo.CategoryName,
	}

	subCategoryInformation := resource.SubCategoryInformation{
		SubCategoryReference: i.SubCategoryInfo.SubCategoryReference,
		SubCategoryName:      i.SubCategoryInfo.SubCategoryName,
	}

	err := r.Client.C("items").Update(bson.M{"reference": reference}, bson.M{"$set": bson.M{
		"item_name":         i.ItemName,
		"description":       i.Description,
		"unit":              i.Unit,
		"location_info":     locationInformation,
		"publish_location":  false,
		"category_info":     categoryInformation,
		"sub_category_info": subCategoryInformation,
		"sponsored":         false,
		"available":         false,
		"updated_at":        time.Now(),
	}})
	if err != nil {
		log.Println("unable to update data:", err.Error())
	}

	return err
}

func (r *MongoRepository) DeleteItem(reference string) error {
	err := r.Client.C("items").Remove(bson.M{"reference": reference})
	if err != nil {
		log.Println("unable to delete data:", err.Error())
	}
	return err
}

func (r *MongoRepository) DeleteAllUserItems(userReference string) error {
	err := r.Client.C("items").Remove(bson.M{"user_reference": userReference})
	if err != nil {
		log.Println("unable to delete data:", err.Error())
	}
	return err
}

func (r *MongoRepository) CheckItemAvailability(reference string, available string) error {
	var item resource.Inventory
	x, _ := strconv.ParseBool(available)
	if x {
		item.Available = true
	} else {
		item.Available = false
	}

	err := r.Client.C("items").Update(bson.M{"reference": reference}, bson.M{"$set": bson.M{
		"available": item.Available,
	}})
	if err != nil {
		log.Println("unable to update availability:", err.Error())
	}
	return err
}

func (r *MongoRepository) AddDocuments(reference string, documents []resource.DocumentInformation) error {
	err := r.Client.C("items").Update(bson.M{"reference": reference}, bson.M{"$set": bson.M{
		"document_info": documents,
	}})
	if err != nil {
		log.Println("unable to update documents:", err.Error())
	}
	return err
}

func (r *MongoRepository) ValidateItem(reference string, validation resource.ValidationInformation) error {
	err := r.Client.C("items").Update(bson.M{"reference": reference}, bson.M{"$set": bson.M{
		"validation_info": validation,
	}})
	if err != nil {
		log.Println("unable to update validation:", err.Error())
	}
	return err
}

func (r *MongoRepository) GetItemDetail(reference string) (*resource.Inventory, error) {
	var item resource.Inventory
	err := r.Client.C("items").Find(bson.M{"reference": reference}).One(&item)
	if err != nil {
		log.Println("unable to get data:", err.Error())
	}
	return &item, err
}

func (r *MongoRepository) GetItemsByMultipleSearch(search string, page int) ([]resource.Inventory, error) {
	var items []resource.Inventory
	err := r.Client.C("items").Find(bson.M{"$or": []bson.M{
		{"item_name": bson.RegEx{search, "i"}},
		{"description": bson.RegEx{search, "i"}},
		{"category_info.category_name": bson.RegEx{search, "i"}},
		{"sub_category_info.sub_category_name": bson.RegEx{search, "i"}},
		{"location_info.state_name": bson.RegEx{search, "i"}},
		{"location_info.city_name": bson.RegEx{search, "i"}},
	}}).Skip((page - 1) * 10).Sort("created_at").Limit(10).All(&items)
	if err != nil {
		log.Println("unable to get data:", err.Error())
	}
	return items, err
}

func (r *MongoRepository) GetItemsByOrganisationSearch(organisationReference string, search string, page int) ([]resource.Inventory, error) {
	var items []resource.Inventory
	err := r.Client.C("items").Find(bson.M{"organization_reference": organisationReference, "$or": []bson.M{
		{"item_name": bson.RegEx{search, "i"}},
		{"description": bson.RegEx{search, "i"}},
		{"category_info.category_name": bson.RegEx{search, "i"}},
		{"sub_category_info.sub_category_name": bson.RegEx{search, "i"}},
		{"location_info.state_name": bson.RegEx{search, "i"}},
		{"location_info.city_name": bson.RegEx{search, "i"}},
	}}).Skip((page - 1) * 10).Sort("created_at").Limit(10).All(&items)
	if err != nil {
		log.Println("unable to get data:", err.Error())
	}
	return items, err
}

func (r *MongoRepository) GetItemsByCategoryAndOrganisationSearch(categoryReference string, organisationReference string, search string, page int) ([]resource.Inventory, error) {
	var items []resource.Inventory
	err := r.Client.C("items").Find(bson.M{"category_info.category_reference": categoryReference, "organization_reference": organisationReference, "$or": []bson.M{
		{"item_name": bson.RegEx{search, "i"}},
		{"description": bson.RegEx{search, "i"}},
		{"category_info.category_name": bson.RegEx{search, "i"}},
		{"sub_category_info.sub_category_name": bson.RegEx{search, "i"}},
		{"location_info.state_name": bson.RegEx{search, "i"}},
		{"location_info.city_name": bson.RegEx{search, "i"}},
	}}).Skip((page - 1) * 5).Sort("created_at").Limit(5).All(&items)
	if err != nil {
		log.Println("unable to get data:", err.Error())
	}
	return items, err
}
