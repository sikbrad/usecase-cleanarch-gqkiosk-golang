package main

import (
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/usecase/customer/ordering"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/dataprovider"
	"log"
)

func init() {
	log.Println("initializing gqskiosk server")

}

func main(){
	log.Println("starting gqskiosk server")

	//uc := ordering.CustomerOrderingUsecase{
	//	GetAllFoodMenuPort: nil,
	//}
	dataprovider := dataprovider.NewCustomerOrderingDbmockDataStorage()
	uc := ordering.NewCustomerOrderingUsecase(dataprovider)
	_ = uc

	foodmenus, _ := uc.GetAllFoodMenu()
	log.Printf("Foodmenus : %#v\n", foodmenus)

	log.Println("finishing gqskiosk server")
}
