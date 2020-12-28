package ordering

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"

type GetAllFoodMenuPorter interface{
	GetAllFoodMenu() ([]entity.FoodMenu, error)
}