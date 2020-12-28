package dataprovider

import (
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
)

var (
	foodMenus	[]entity.FoodMenu
)

type repo struct{

}

func (_ *repo) GetFoodMenuWithName(searchName string) (*entity.FoodMenu) {

	var foodMenu *entity.FoodMenu
	for _, v := range foodMenus {
		if v.Name == searchName {
			foodMenu = &v
		}
	}

	return foodMenu
}

func (_ *repo) GetAllFoodMenu() (*[]entity.FoodMenu) {
	return &foodMenus
}

func NewCustomerOrderingDbmockDataStorage() DataProvider{

	// for debugging, add data
	foodMenus = append(foodMenus, *entity.NewFoodMenu(
		"Big Mac",
		"McDonalds signature burger",
		"http://google.com/image.png"),
	)
	foodMenus = append(foodMenus, *entity.NewFoodMenu(
		"Cheese Burger",
		"Budget choice",
		"http://google.com/image.png"),
	)

	return &repo{

	}
}
