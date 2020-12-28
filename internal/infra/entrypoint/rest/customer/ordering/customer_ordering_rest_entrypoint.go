package ordering

import (
	"encoding/json"
	"fmt"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/core/usecase/customer/ordering"
	"github.com/sikbrad/usecase-cleanarch-gqkiosk-golang/internal/infra/entrypoint/rest/errors"
	"net/http"
	"net/url"
	"path"
)

type CustomerOrderingRestEntrypoint interface {
	GetAllFoodMenus(resp http.ResponseWriter, request *http.Request)
	GetFoodMenuWithName(resp http.ResponseWriter, request *http.Request)
}

type customerOrderingRestEntrypoint struct {
	usecase *ordering.CustomerOrderingUsecase
}

func NewCustomerOrderingRestEntrypoint(usecase *ordering.CustomerOrderingUsecase) CustomerOrderingRestEntrypoint{
	return &customerOrderingRestEntrypoint{
		usecase: usecase,
	}
}

func (c *customerOrderingRestEntrypoint) GetAllFoodMenus(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	foodMenus, err := c.usecase.GetAllFoodMenu()

	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "Error when parsing message",
		})

		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(foodMenus)
}

func (c *customerOrderingRestEntrypoint) GetFoodMenuWithName(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	escaped_name := path.Base(request.URL.Path)
	//escaped_name := request.URL.Query().Get("name") ## can't use it -> maybe abstracting router was overkill. it conflicts.
	fmt.Println("escaped_name : ",escaped_name)
	name, _ := url.QueryUnescape(escaped_name)
	fmt.Println("name : ",name)

	// no param
	if len(name) == 0 {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "Error accepting url parameter",
		})

		return
	}

	foodMenu, err := c.usecase.GetFoodMenuWithName(name)

	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{
			Message : "Error when parsing message",
		})

		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(foodMenu)
}