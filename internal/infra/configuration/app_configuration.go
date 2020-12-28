package configuration

import (
	"fmt"
	usecase_customer_ordering "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/usecase/customer/ordering"
	entrypoint_customer_ordering "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/entrypoint/rest/customer/ordering"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/dataprovider"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/entrypoint/router"
	"log"
	"net/http"
)

const (
	port string = ":8000"
)

func Init(){

	/*
	 database config
	 */
	//either can work
	//dataprovider := dataprovider.NewCustomerOrderingDbmockDataStorage() //uses list as db
	dataprovider := dataprovider.NewCustomerOrderingSqliteDataStorage() //uses sqlite db

	uc := usecase_customer_ordering.NewCustomerOrderingUsecase(
		dataprovider, //getAllFoodMenuPorter
		dataprovider, //getFoodMenuWithNamePorter
	)

	//prints for testing
	foodmenus, _ := uc.GetAllFoodMenu()
	log.Printf("Getting all Foodmenus : %#v\n", foodmenus)
	foodmenu, _ := uc.GetFoodMenuWithName("Big Mac")
	log.Printf("Get foodmenu with name 'Big Mac' : %#v\n", foodmenu)

	/*
	router config
	 */
	httpRouter := router.NewChiRouter() //chi as router
	//httpRouter := router.NewMuxRouter() //mux as router

	// Router
	customerOrderingRestEntrypoint := entrypoint_customer_ordering.NewCustomerOrderingRestEntrypoint(uc)

	httpRouter.GET("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp,"entered index")
	})
	httpRouter.GET("/api/customer/ordering/foodmenus", customerOrderingRestEntrypoint.GetAllFoodMenus)
	httpRouter.GET("/api/customer/ordering/foodmenu/{name}", customerOrderingRestEntrypoint.GetFoodMenuWithName)
	httpRouter.SERVE(port)

	//to test
	// access http://localhost:8000/api/customer/ordering/foodmenus
	// access http://localhost:8000/api/customer/ordering/foodmenu/Big%20mac

}