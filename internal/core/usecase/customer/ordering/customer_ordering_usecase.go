package ordering

import (
	"errors"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core"
)

type CustomerOrderingUsecase struct {
	getAllFoodMenuPorter GetAllFoodMenuPorter
	getFoodMenuWithNamePorter GetFoodMenuWithNamePorter
}

func NewCustomerOrderingUsecase(
	getAllFoodMenuPorter GetAllFoodMenuPorter,
	getFoodMenuWithNamePorter GetFoodMenuWithNamePorter,
	) *CustomerOrderingUsecase{

	return &CustomerOrderingUsecase{
		getAllFoodMenuPorter: getAllFoodMenuPorter,
		getFoodMenuWithNamePorter: getFoodMenuWithNamePorter,
	}
}

func (uc *CustomerOrderingUsecase) GetAllFoodMenu() (*[]entity.FoodMenu, error) {
	menus := uc.getAllFoodMenuPorter.GetAllFoodMenu()

	if len(*menus) == 0{
		return nil, errors.New(core.ERROR_NO_FOODMENU_WITH_GIVEN_NAME)
	}

	return menus, nil
}

func (uc *CustomerOrderingUsecase) GetFoodMenuWithName(foodName string) (*entity.FoodMenu, error) {
	menu := uc.getFoodMenuWithNamePorter.GetFoodMenuWithName(foodName)

	if menu == nil {
		return nil, errors.New(core.ERROR_NO_FOODMENU_WITH_GIVEN_NAME)
	}

	return menu, nil
}