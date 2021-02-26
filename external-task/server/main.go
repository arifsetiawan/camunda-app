package main

import (
	"sync"

	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"

	"github.com/arifsetiawan/camunda-app/external-task/handler"
	"github.com/arifsetiawan/camunda-app/pkg/camunda"
	"github.com/arifsetiawan/camunda-app/pkg/env"
)

func main() {
	camundaAPIURL := env.Getenv("CAMUNDA_API_URL", "http://localhost:8080/engine-rest")
	camundaClient := camunda.NewClient(camundaAPIURL)

	availableLeaveDays := &camunda.FetchAndLockRequest{
		WorkerID: "available-leave-days-worker",
		MaxTasks: 1,
		Topics: []camunda.FetchAndLockTopic{
			camunda.FetchAndLockTopic{
				Name:         "available-leave-days",
				LockDuration: 10000,
			},
		},
	}

	email := &camunda.FetchAndLockRequest{
		WorkerID: "email-worker",
		MaxTasks: 1,
		Topics: []camunda.FetchAndLockTopic{
			camunda.FetchAndLockTopic{
				Name:         "send-rejection-email",
				LockDuration: 10000,
			},
			camunda.FetchAndLockTopic{
				Name:         "send-approval-email",
				LockDuration: 10000,
			},
		},
	}

	var wg sync.WaitGroup
	c := cron.New()
	c.AddFunc("@every 10s", call(camundaClient, availableLeaveDays, handler.AvailableLeaveDaysHandler))
	c.AddFunc("@every 10s", call(camundaClient, email, handler.EmailHandler))
	c.Start()

	wg.Add(1)
	wg.Wait()
}

func call(client *camunda.Client, param *camunda.FetchAndLockRequest, handler func(client *camunda.Client, tasks []camunda.ExternalTask) error) func() {
	return func() {
		tasks, err := client.FetchAndLockExternalTask(param)
		if err != nil {
			log.Error().Err(err).Msg("fetch and lock task")
			return
		}

		log.Info().Interface("topics", param.Topics).Interface("tasks", tasks).Msg("Fetch and lock")

		if len(*tasks) > 0 {
			err = handler(client, *tasks)
			if err != nil {
				log.Error().Err(err).Msg("task handler")
				return
			}
		}
	}
}
