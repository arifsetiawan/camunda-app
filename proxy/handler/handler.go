// Copyright (c) Mindtrex 2019 All Rights Reserved.

package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
	"github.com/arifsetiawan/camunda-app/pkg/env"
)

// InitHandler ...
func InitHandler(e *echo.Echo) {
	camundaAPIURL := env.Getenv("CAMUNDA_API_URL", "http://localhost:8080/engine-rest")
	camundaClient := camunda.NewClient(camundaAPIURL)

	// Auth
	a := e.Group("/auth")

	authHandler := AuthHandler{
		CamundaClient: camundaClient,
	}
	authHandler.SetRoutes(a)

	// Routes with JWT authentication
	r := e.Group("")
	config := middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	taskHandler := TaskHandler{
		CamundaClient: camundaClient,
	}
	taskHandler.SetRoutes(r)
}
