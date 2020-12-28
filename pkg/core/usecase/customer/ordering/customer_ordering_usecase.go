package ordering

import (
	"../../../entity"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/core/entity"
)

var (
	fm entity.FoodMenu
)

type CustomerOrderingUsecase interface {
	GetAllFoodMenu() []entity.FoodMenu
}
