package ordering

import (
	"../../../entity"
)


type CustomerOrderingUsecase struct {
	GetAllFoodMenuPort GetAllFoodMenuPorter
}

func NewCustomerOrderingUsecase(getAllFoodMenuPort GetAllFoodMenuPorter) *CustomerOrderingUsecase{
	return &CustomerOrderingUsecase{
		GetAllFoodMenuPort: getAllFoodMenuPort,
	}
}

func (uc *CustomerOrderingUsecase) GetAll() []entity.FoodMenu {
	return uc.GetAllFoodMenuPort.GetAll()
}