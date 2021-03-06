package main

import (
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/configuration"
	"log"
)

func init() {
	log.Println("initializing gqskiosk server")

}

func main() {
	log.Println("starting gqskiosk server")

	configuration.Init()

	log.Println("finishing gqskiosk server")
}

//func run1() {
//	dataprovider := dataprovider.NewCustomerOrderingDbmockDataStorage()
//	uc := ordering.NewCustomerOrderingUsecase(
//		dataprovider, //getAllFoodMenuPorter
//		dataprovider, //getFoodMenuWithNamePorter
//	)
//
//	foodmenus, _ := uc.GetAllFoodMenu()
//	log.Printf("Getting all Foodmenus : %#v\n", foodmenus)
//
//	foodmenu, _ := uc.GetFoodMenuWithName("Big Mac")
//	log.Printf("Get foodmenu with name 'Big Mac' : %#v\n", foodmenu)
//}
//
//func run2() {
//	dataprovider := dataprovider.NewCustomerOrderingSqliteDataStorage()
//	uc := ordering.NewCustomerOrderingUsecase(
//		dataprovider, //getAllFoodMenuPorter
//		dataprovider, //getFoodMenuWithNamePorter
//	)
//
//	foodmenus, _ := uc.GetAllFoodMenu()
//	log.Printf("Getting all Foodmenus : %#v\n", foodmenus)
//
//	foodmenus, _ = uc.GetAllFoodMenu()
//	log.Printf("Getting all Foodmenus2 : %#v\n", foodmenus)
//
//	foodmenu, _ := uc.GetFoodMenuWithName("Burger")
//	log.Printf("Get foodmenu with name 'Big Mac' : %#v\n", foodmenu)
//}
