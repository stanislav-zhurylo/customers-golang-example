package controller

import (
	"customers-example/model"
	"customers-example/service"
	"customers-example/utils"
	"fmt"
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type CustomerControllerInterface interface {
	CreateCustomer(c *gin.Context)
	EditCustomer(c *gin.Context)
	DeleteCustomerById(c *gin.Context)
	RenderCustomerView(c *gin.Context, mode string)
	FindCustomers(c *gin.Context)
}

type CustomerController struct {
	customerService service.CustomerServiceInterface
}

func NewCustomerController(customerService service.CustomerServiceInterface) CustomerControllerInterface {
	return &CustomerController{
		customerService,
	}
}

func (customerController *CustomerController) CreateCustomer(c *gin.Context) {
	var payload model.Customer

	if err := c.Bind(&payload); err != nil {
		var ve validator.ValidationErrors
		errors.As(err, &ve)

		if errors.As(err, &ve) {
			out := make([]model.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = model.ErrorMsg{
					Field:   fe.Field(),
					Message: utils.GetErrorMsg(fe),
				}
			}
			c.HTML(http.StatusOK, "customer.html", gin.H{
				"customer": payload,
				"errors":   out,
				"mode":     "create",
			})
		}
		return

	} else {
		created, err := customerController.customerService.CreateCustomer(c, payload)
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/customers/%d", created.Id))
	}
}

func (customerController *CustomerController) EditCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var payload model.Customer
	payload.Id = int64(id)

	if err := c.Bind(&payload); err != nil {
		var ve validator.ValidationErrors
		errors.As(err, &ve)

		if errors.As(err, &ve) {
			out := make([]model.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = model.ErrorMsg{
					Field:   fe.Field(),
					Message: utils.GetErrorMsg(fe),
				}
			}
			c.HTML(http.StatusOK, "customer.html", gin.H{
				"customer": payload,
				"errors":   out,
				"mode":     "edit",
			})
		}
		return

	} else {
		customerController.customerService.UpdateCustomer(c, payload)
		c.Redirect(http.StatusFound, fmt.Sprintf("/customers/%d", id))
	}
}

func (customerController *CustomerController) RenderCustomerView(c *gin.Context, mode string) {
	if mode == "create" {
		var customer model.Customer
		c.HTML(http.StatusOK, "customer.html", gin.H{
			"customer": customer,
			"mode":     "create",
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := customerController.customerService.GetCustomer(c, id)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "customer.html", gin.H{
		"customer": customer,
		"mode":     mode,
	})
}

func (customerController *CustomerController) DeleteCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err1 := customerController.customerService.GetCustomer(c, id)
	if err1 != nil {
		panic(err1)
	}

	err2 := customerController.customerService.DeleteCustomerById(c, int64(id))
	if err2 != nil {
		panic(err2)
	}

	c.HTML(http.StatusOK, "customer-delete.html", gin.H{
		"customer": customer,
	})
}

func (customerController *CustomerController) FindCustomers(c *gin.Context) {

	var searchFilter model.SearchFilter
	var paginationSettings model.PaginationSettings
	decoder.Decode(&searchFilter, c.Request.URL.Query())
	decoder.Decode(&paginationSettings, c.Request.URL.Query())

	if paginationSettings.Page == 0 {
		paginationSettings.Page = 1
	}
	if paginationSettings.PageSize == 0 {
		paginationSettings.PageSize = 10
	}
	if paginationSettings.Sort == "" {
		paginationSettings.Sort = "id"
	}
	if paginationSettings.Direction == "" {
		paginationSettings.Direction = "asc"
	}

	customers, totalCount, err := customerController.customerService.FindCustomers(c, paginationSettings, searchFilter)
	if err != nil {
		panic(err)
	}

	totalPages := 1
	if totalCount%paginationSettings.PageSize != 0 {
		totalPages = totalCount/paginationSettings.PageSize + 1
	} else {
		totalPages = totalCount / paginationSettings.PageSize
	}

	c.HTML(http.StatusOK, "customers.html", gin.H{
		"customers":    customers,
		"totalCount":   totalCount,
		"pages":        utils.MakeRange(1, totalPages),
		"pagesCount":   totalPages,
		"pagination":   paginationSettings,
		"searchFilter": searchFilter,
	})
}
