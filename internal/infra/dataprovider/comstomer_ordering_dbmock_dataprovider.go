package dataprovider

import (
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/entity"
)

var (
	foodMenus	[]entity.FoodMenu
)

type repo struct{

}

func (r *repo) GetAll() ([]entity.FoodMenu, error) {
	return foodMenus, nil
}

//type DbmockDataProvider interface{
//	DataProvider
//}


func NewCustomerOrderingDbmockDataStorage() DataProvider{

	// for debugging, add data
	foodMenus = append(foodMenus, entity.FoodMenu{
		Id:          "some_uuid",
		Name:        "Big Mac",
		Description: "McDonalds signature burger",
		ImageUrl:    "http://google.com/image.png",
	})

	return &repo{

	}
}

//func (dataprovider* DbmockDataProvider) FindAll() ([]entity.FoodMenu, error) {
//	//fetch datastore and provide data
//	//in our case, just get array of foodmenus list of all.
//	return foodMenus, nil
//}
//

//
//func NewGetAllFoodMenuPorter() ordering.GetAllFoodMenuPorter{
//	return &DbmockDataProvider{
//
//	}
//	//return func() ([]entity.FoodMenu, error){
//	//
//	//}
//}
//func (port *ordering.GetAllFoodMenuPorter) GetAll() ([]entity.FoodMenu, error) {
//	return uc.getAllFoodMenuPort.GetAll()
//}
