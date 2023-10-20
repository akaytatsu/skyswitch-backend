package main

import (
	"app/api"
	"app/cron"
	"app/infrastructure/postgres"
)

func main() {
	cron.StartCronJobs()

	postgres.Migrations()

	api.StartWebServer()
}
