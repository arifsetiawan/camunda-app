package handler

import (
	"time"

	"github.com/arifsetiawan/camunda-app/pkg/camunda"
)

// EmailHandler ...
func EmailHandler(client *camunda.Client, tasks []camunda.ExternalTask) error {
	for _, task := range tasks {

		// pretend to do something
		time.Sleep(2 * time.Second)

		// complete
		err := client.CompleteExternalTask(task.ID, &camunda.CompleteExternalTaskRequest{
			WorkerID: "email-worker",
		})
		if err != nil {
			return err
		}
	}

	return nil
}
