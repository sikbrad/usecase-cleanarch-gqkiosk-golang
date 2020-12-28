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

	dataprovider := dataprovider.NewCustomerOrderingDbmockDataStorage()
	uc := ordering.NewCustomerOrderingUsecase(
		dataprovider, //getAllFoodMenuPorter
		dataprovider, //getFoodMenuWithNamePorter
		)

	foodmenus, _ := uc.GetAllFoodMenu()
	log.Printf("Getting all Foodmenus : %#v\n", foodmenus)

	foodmenu, _ := uc.GetFoodMenuWithName("Big Mac")
	log.Printf("Get foodmenu with name 'Big Mac' : %#v\n", foodmenu)

	log.Println("finishing gqskiosk server")
}
