package dataprovider

import (
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
)

var (
	foodMenus	[]entity.FoodMenu
)

type repo struct{

}

func (r *repo) GetAllFoodMenu() ([]entity.FoodMenu, error) {
	return foodMenus, nil
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
