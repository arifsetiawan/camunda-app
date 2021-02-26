package handler

import (
	"time"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// AvailableLeaveDaysHandler ...
func AvailableLeaveDaysHandler(client *camunda.Client, tasks []camunda.ExternalTask) error {
	for _, task := range tasks {

		// pretend to do something
		time.Sleep(2 * time.Second)

		// complete
		err := client.CompleteExternalTask(task.ID, &camunda.CompleteExternalTaskRequest{
			WorkerID: "worker1",
			Variables: map[string]camunda.Variable{
				"daysAvailable": camunda.Variable{
					Value: true,
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
