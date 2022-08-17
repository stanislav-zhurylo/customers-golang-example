package main

import (
	"customers-example/controller"
	"customers-example/migrations"
	"customers-example/repository"
	"customers-example/router"
	"customers-example/service"
	"customers-example/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	pgKit "github.com/laironacosta/kit-go/postgresql"
	"github.com/pkg/errors"
)

var DatabaseConfiguration struct {
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASS" default:"changeme"`
	Database string `envconfig:"DB_NAME" default:"clog-local"`
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     int    `envconfig:"DB_PORT" default:"5432"`
}

func main() {
	gin := gin.Default()

	if err := envconfig.Process("LIST", &DatabaseConfiguration); err != nil {
		errors.Wrap(err, "parse environment variables")
		return
	}

	db := pgKit.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", DatabaseConfiguration.Host, DatabaseConfiguration.Port),
		User:     DatabaseConfiguration.User,
		Password: DatabaseConfiguration.Password,
		Database: DatabaseConfiguration.Database,
	})

	migrations.Init(db)

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	utils.LoadTemplates("views/*.html")

	router.NewRouter(gin, customerController).Init()

	// listen and serve on 0.0.0.0:8080
	gin.Run()
}
