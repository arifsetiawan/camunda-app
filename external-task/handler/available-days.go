package handler

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// AvailableLeaveDaysHandler ...
func AvailableLeaveDaysHandler(client *camunda.Client, tasks []camunda.ExternalTask) error {
	for _, task := range tasks {

		// pretend to do something
		time.Sleep(2 * time.Second)

		log.Info().Interface("task", task).Msg("task")

		// complete
		err := client.CompleteExternalTask(task.ID, &camunda.CompleteExternalTaskRequest{
			WorkerID: "available-leave-days-worker",
			Variables: map[string]camunda.Variable{
				"daysAvailable": camunda.Variable{
					Value: true,
					Type:  "Boolean",
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
