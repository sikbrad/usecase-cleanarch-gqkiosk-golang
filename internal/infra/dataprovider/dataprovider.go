package dataprovider

import "github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/usecase/customer/ordering"

type DataProvider interface {
	ordering.GetAllFoodMenuPorter
}


