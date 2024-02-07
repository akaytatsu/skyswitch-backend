package main

import (
	"app/api"
	"app/config"
	"app/cron"
	"app/infrastructure/postgres"
)

func main() {
	config.ReadEnvironmentVars()

	cron.StartCronJobs()

	// conn := postgres.Connect()
	postgres.Migrations()

	// usecase := usecase_user.NewService(
	// 	repository.NewUserPostgres(conn),
	// )

	// err := usecase.CreateAdminUser()
	// if err != nil {
	// 	log.Println("---------->     Error creating admin user     <----------")
	// 	log.Println(err)
	// }

	api.StartWebServer()
}
