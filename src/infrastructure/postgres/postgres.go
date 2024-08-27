package postgres

import (
	"app/config"
	"app/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func Connect() *gorm.DB {

	if gormDB == nil {
		return conn()
	}

	return gormDB
}

func Migrations() {
	db := Connect()

	db.AutoMigrate(&entity.EntityUser{})
	db.AutoMigrate(&entity.EntityDbinstance{})
	db.AutoMigrate(&entity.EntityCloudAccount{})
	db.AutoMigrate(&entity.EntityInstance{})
	db.AutoMigrate(&entity.EntityHoliday{})
	db.AutoMigrate(&entity.EntityCalendar{})
	db.AutoMigrate(&entity.EntityLog{})
	db.AutoMigrate(&entity.EntityAutoScalingGroup{})
}

func conn() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.EnvironmentVariables.POSTGRES_HOST,
		config.EnvironmentVariables.POSTGRES_USER,
		config.EnvironmentVariables.POSTGRES_PASSWORD,
		config.EnvironmentVariables.POSTGRES_DB,
		config.EnvironmentVariables.POSTGRES_PORT,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	gormDB = conn

	return gormDB
}
