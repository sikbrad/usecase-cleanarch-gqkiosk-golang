package entity

import "github.com/google/uuid"

type FoodMenu struct{
	Id string			`json:"id"`
	Name string			`json:"name"`
	Description string	`json:"description"`
	ImageUrl string		`json:"image_url"`
}

func NewFoodMenu(name, description, imageUrl string) *FoodMenu{
	newUuid, _ := uuid.NewRandom()
	newUuidStr := newUuid.String()

	return &FoodMenu{
		Id:          newUuidStr,
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
	}
}