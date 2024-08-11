package api

import (
	"log"

	"app/api/handlers"
	"app/config"
	"app/cron"
	infrastructure_cloud_provider_aws "app/infrastructure/cloud_provider/aws"
	"app/infrastructure/postgres"
	"app/infrastructure/repository"
	usecase_calendar "app/usecase/calendar"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_dbinstance "app/usecase/dbinstance"
	usecase_holiday "app/usecase/holiday"
	usecase_instance "app/usecase/instance"
	usecase_job "app/usecase/job"
	usecase_log "app/usecase/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupDatabase() *gorm.DB {
	conn := postgres.Connect()
	return conn
}

func setupRouter(conn *gorm.DB) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	var usecaseDbinstance usecase_dbinstance.IUsecaseDbinstance = usecase_dbinstance.NewService(
		repository.NewDbinstancePostgres(conn),
	)

	var usecaseLog usecase_log.IUsecaseLog = usecase_log.NewService(
		repository.NewLogPostgres(conn),
	)

	var usecaseInstance usecase_instance.IUseCaseInstance = usecase_instance.NewService(
		repository.NewInstancePostgres(conn),
	)

	var usecaseCloudAccount usecase_cloud_account.IUsecaseCloudAccount = usecase_cloud_account.NewAWSService(
		repository.NewCloudAccountPostgres(conn),
		usecaseInstance,
		infrastructure_cloud_provider_aws.NewAWSCloudProvider(),
		usecaseDbinstance,
	)

	var usecaseHoliday usecase_holiday.IUsecaseHoliday = usecase_holiday.NewService(
		repository.NewHolidayPostgres(conn),
	)

	var usecaseCalendar usecase_calendar.IUsecaseCalendar = usecase_calendar.NewService(
		repository.NewCalendarPostgres(conn),
		cron.Scheduler,
		usecaseInstance,
		infrastructure_cloud_provider_aws.NewAWSCloudProvider(),
		usecaseCloudAccount,
		usecaseHoliday,
		usecaseLog,
	)

	var usecaseJob usecase_job.IUsecaseJob = usecase_job.NewService(
		usecaseCalendar,
	)

	handlers.MountCloudAccountHandlers(r, conn)
	handlers.MountUsersHandlers(r, conn)
	handlers.MountDbinstanceRoutes(r, conn, usecaseDbinstance)
	handlers.MountInstancesRoutes(r, conn)
	handlers.MountCalendarRoutes(r, conn, usecaseCalendar)
	handlers.MountJobRoutes(r, conn, usecaseJob)
	handlers.MountLogRoutes(r, conn, usecaseLog)

	return r
}

func SetupRouters() *gin.Engine {
	conn := setupDatabase()
	return setupRouter(conn)
}

func StartWebServer() {
	config.ReadEnvironmentVars()

	r := SetupRouters()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// se for release, reduz o log
	if config.EnvironmentVariables.ISRELEASE {
		gin.SetMode(gin.ReleaseMode)
	}

	// Bind to a port and pass our router in
	log.Fatal(r.Run())
}
