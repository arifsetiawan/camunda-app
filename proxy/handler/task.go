package handler

import (
	"errors"
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
	e.GET("/tasks/:id", h.taskDetail)
	e.POST("/tasks/:id/complete", h.completeTask)
}

func (h *TaskHandler) listTask(c echo.Context) (err error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*jwt.StandardClaims)

	userGroups, err := h.CamundaClient.IdentityGroups(claims.Subject)
	if err != nil {
		return err
	}

	allTasks := make([]camunda.UserTask, 0)
	for _, group := range userGroups.Groups {
		tasks, err := h.CamundaClient.ListUserTask(&camunda.ListUserTaskRequest{
			CandidateGroup: group.ID,
			Unassigned:     true,
		})
		if err != nil {
			return err
		}

		allTasks = append(allTasks, *tasks...)
	}

	return c.JSON(http.StatusOK, allTasks)
}

func (h *TaskHandler) taskDetail(c echo.Context) (err error) {
	taskID := c.Param("id")

	variables, err := h.CamundaClient.UserTaskVariables(taskID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, variables)
}

func (h *TaskHandler) completeTask(c echo.Context) (err error) {
	taskID := c.Param("id")

	data := new(camunda.CompleteUserTaskRequest)
	err = c.Bind(data)
	if err != nil {
		return errors.New("Failed to get process start data. Probably content-type is not match with actual body type")
	}

	err = h.CamundaClient.CompleteUserTask(taskID, data)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
