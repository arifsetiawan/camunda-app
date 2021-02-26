package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
	"github.com/arifsetiawan/camunda-app/pkg/middleware"
)

// TaskHandler ...
type TaskHandler struct {
	CamundaClient *camunda.Client
}

// SetRoutes ...
func (h *TaskHandler) SetRoutes(e *echo.Group) {
	e.GET("/tasks", h.listTask, middleware.ListQueryParams)
}

func (h *TaskHandler) listTask(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	// verify user to camunda
	tasks, err := h.CamundaClient.ListTask(&camunda.ListTaskRequest{
		Assignee: claims.Subject,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}
