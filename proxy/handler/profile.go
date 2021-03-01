package handler

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// ProfileHandler ...
type ProfileHandler struct {
	CamundaClient *camunda.Client
}

// EmployeeData ...
type EmployeeData struct {
	RoleID    string `json:"roleId"`
	RoleName  string `json:"roleName"`
	LeaveDays int    `json:"leaveDays"`
}

var employees = map[string]EmployeeData{
	"sampurple": EmployeeData{
		RoleID:    "juniors",
		RoleName:  "Junior Engineer",
		LeaveDays: 15,
	},
	"annepink": EmployeeData{
		RoleID:    "seniors",
		RoleName:  "Senior Engineer",
		LeaveDays: 20,
	},
	"johnblack": EmployeeData{
		RoleID:    "hr",
		RoleName:  "Human Resources",
		LeaveDays: 20,
	},
	"sophiagreen": EmployeeData{
		RoleID:    "manager",
		RoleName:  "Managers",
		LeaveDays: 30,
	},
	"markwhite": EmployeeData{
		RoleID:    "ceo",
		RoleName:  "CEO",
		LeaveDays: 30,
	},
}

// SetRoutes ...
func (h *ProfileHandler) SetRoutes(e *echo.Group) {
	e.GET("/me", h.me)
	e.GET("/employee", h.employee)
	e.GET("/groups", h.groups)
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

func (h *ProfileHandler) employee(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	data, ok := employees[claims.Subject]
	if !ok {
		return fmt.Errorf("employee data is not found")
	}

	return c.JSON(http.StatusOK, data)
}

func (h *ProfileHandler) groups(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	profile, err := h.CamundaClient.IdentityGroups(claims.Subject)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profile)
}
