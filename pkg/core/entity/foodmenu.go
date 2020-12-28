package entity

type FoodMenu struct{
	Id string			`json:"id"`
	Name string			`json:"name"`
	Description string	`json:"description"`
	ImageUrl string		`json:"image_url"`
}