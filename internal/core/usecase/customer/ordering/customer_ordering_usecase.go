package ordering

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"

type CustomerOrderingUsecase struct {
	getAllFoodMenuPort GetAllFoodMenuPorter
}

func NewCustomerOrderingUsecase(
	getAllFoodMenuPort GetAllFoodMenuPorter,
	) *CustomerOrderingUsecase{

	return &CustomerOrderingUsecase{
		getAllFoodMenuPort: getAllFoodMenuPort,
	}
}

func (uc *CustomerOrderingUsecase) GetAllFoodMenu() ([]entity.FoodMenu, error) {
	return uc.getAllFoodMenuPort.GetAll()
}
