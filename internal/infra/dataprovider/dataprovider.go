package dataprovider

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/usecase/customer/ordering"

//listing interface will delegate the implementation of porters into this interface.
type DataProvider interface {
	ordering.GetAllFoodMenuPorter
	ordering.GetFoodMenuWithNamePorter
}


