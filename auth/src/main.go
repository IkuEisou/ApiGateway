//go:generate goagen bootstrap -d github.com/ikueisou/apigateway/auth/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ikueisou/apigateway/auth/src/app"
	"github.com/ikueisou/apigateway/auth/utils/database"
	"os"
)

func main() {
	// Create service
	service := goa.New("Authentication API")

	db, err := database.Connect(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"))
	if err != nil {
		service.LogError("startup", "err", err)
	}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "User" controller
	c := NewUserController(service, db)
	app.MountUserController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
