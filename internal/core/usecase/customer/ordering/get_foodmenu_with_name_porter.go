package ordering

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"

type GetFoodMenuWithNamePorter interface{
	GetFoodMenuWithName(searchName string) (*entity.FoodMenu)
}