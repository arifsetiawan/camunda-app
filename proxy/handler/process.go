package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
	"github.com/arifsetiawan/camunda-app/pkg/random"
)

// ProcessHandler ...
type ProcessHandler struct {
	CamundaClient *camunda.Client
}

// SetRoutes ...
func (h *ProcessHandler) SetRoutes(e *echo.Group) {
	e.POST("/process/start", h.start)
}

func (h *ProcessHandler) start(c echo.Context) (err error) {
	//token := c.Get("user").(*jwt.Token)
	//claims := token.Claims.(*jwt.StandardClaims)

	data := new(camunda.ProcessStartRequest)
	err = c.Bind(data)
	if err != nil {
		return errors.New("Failed to get process start data. Probably content-type is not match with actual body type")
	}

	log.Info().Interface("data", data).Msg("start process 1")

	data.BusinessKey = random.GenerateAlphaNumeric(16)
	data.Variables["leaveId"] = camunda.Variable{
		Type:  "String",
		Value: data.BusinessKey,
	}

	log.Info().Interface("data", data).Msg("start process 2")

	err = h.CamundaClient.ProcessDefinitionStart("Process_1bvk6g2", data)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
