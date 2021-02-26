package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// AuthHandler ...
type AuthHandler struct {
	CamundaClient *camunda.Client
}

// Login ...
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SetRoutes ...
func (h *AuthHandler) SetRoutes(e *echo.Group) {
	e.POST("/login", h.login)
}

func (h *AuthHandler) login(c echo.Context) (err error) {
	login := new(Login)
	err = c.Bind(login)
	if err != nil {
		return errors.New("Failed to get login data. Probably content-type is not match with actual body type")
	}

	// verify user to camunda
	user, err := h.CamundaClient.IdentityVerify(&camunda.IdentityVerifyRequest{
		Username: login.Username,
		Password: login.Password,
	})
	if err != nil {
		return err
	}

	if !user.Authenticated {
		return c.JSON(http.StatusUnauthorized, user)
	}

	// Set custom claims
	claims := &jwt.StandardClaims{
		Audience:  "Camunda Custom Tasklist",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "Camunda Proxy",
		Subject:   user.AuthenticatedUser,
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
