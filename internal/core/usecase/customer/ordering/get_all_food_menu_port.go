package ordering

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"

//as golang recommends to youse 'getter' in their interface name, i changed it.

type GetAllFoodMenuPorter interface{
	GetAll() ([]entity.FoodMenu, error)
}