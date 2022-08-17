package router

import (
	"customers-example/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	server             *gin.Engine
	customerController controller.CustomerControllerInterface
}

func NewRouter(
	server *gin.Engine,
	customerController controller.CustomerControllerInterface) *Router {

	return &Router{
		server,
		customerController,
	}
}

func (router *Router) Init() {
	router.server.Static("/assets", "./assets")
	router.server.LoadHTMLGlob("views/*.html")

	rootGroup := router.server.Group("/")
	rootGroup.GET("/customers", router.customerController.FindCustomers)
	rootGroup.POST("/customers", router.customerController.CreateCustomer)
	rootGroup.POST("/customers/:id", router.customerController.EditCustomer)
	rootGroup.GET("/customers/:id/delete", router.customerController.DeleteCustomerById)

	rootGroup.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/customers")
	})
	rootGroup.GET("/customers/:id", func(c *gin.Context) {
		router.customerController.RenderCustomerView(c, "view")
	})
	rootGroup.GET("/customers/:id/edit", func(c *gin.Context) {
		router.customerController.RenderCustomerView(c, "edit")
	})
	rootGroup.GET("/customers/create", func(c *gin.Context) {
		router.customerController.RenderCustomerView(c, "create")
	})
}
