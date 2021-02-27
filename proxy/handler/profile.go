package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// ProfileHandler ...
type ProfileHandler struct {
	CamundaClient *camunda.Client
}

// SetRoutes ...
func (h *ProfileHandler) SetRoutes(e *echo.Group) {
	e.GET("/me", h.me)
}

func (h *ProfileHandler) me(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	profile, err := h.CamundaClient.UserProfile(claims.Subject)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profile)
}
