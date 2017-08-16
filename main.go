//go:generate goagen bootstrap -d github.com/JormungandrK/jwt-issuer/design

package main

import (
	"os"

	"github.com/JormungandrK/jwt-issuer/api"
	"github.com/JormungandrK/jwt-issuer/app"
	"github.com/JormungandrK/jwt-issuer/config"
	"github.com/JormungandrK/jwt-issuer/store"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	cf := os.Getenv("SERVICE_CONFIG_FILE")
	if cf == "" {
		cf = "config.json"
	}
	config, err := config.LoadConfig(cf)
	if err != nil {
		panic(err)
	}

	keyStore, err := store.NewFileKeyStore(config.Keys)
	if err != nil {
		panic(err)
	}

	userAPI, err := api.NewUserAPI(config, keyStore)
	if err != nil {
		panic(err)
	}
	// Create service
	service := goa.New("jwt-signin")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "signin" controller
	c := NewSigninController(service, userAPI, keyStore, config)
	app.MountJWTController(service, c)

	// Start service
	if err := service.ListenAndServe(":8085"); err != nil {
		service.LogError("startup", "err", err)
	}

}
