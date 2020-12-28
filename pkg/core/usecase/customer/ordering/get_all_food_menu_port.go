package ordering

import (
	"../../../entity"
)

//as golang recommends to youse 'getter' in their interface name, i changed it.

type GetAllFoodMenuPorter interface{
	GetAll() []entity.FoodMenu
}