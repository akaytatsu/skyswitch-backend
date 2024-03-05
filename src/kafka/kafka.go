package kafka

import (
	"app/config"
	"app/entity"
	"app/infrastructure/postgres"
	"app/infrastructure/repository"
	usecase_holiday "app/usecase/holiday"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func StartKafka() {

	db := postgres.Connect()

	repoHoliday := repository.NewHolidayPostgres(db)
	usecaseHoliday := usecase_holiday.NewService(repoHoliday)

	var topicParams []KafkaReadTopicsParams

	topicParams = append(topicParams, KafkaReadTopicsParams{
		Topic: config.EnvironmentVariables.KAFKA_HOLIDAY_TOPIC,
		Handler: func(m kafka.Message) error {
			var holiday []string

			err := json.Unmarshal(m.Value, &holiday)

			if err != nil {
				return err
			}

			for _, dateStr := range holiday {
				date, err := usecaseHoliday.DateStringToTime(dateStr)

				if err != nil {
					return err
				}

				err = usecaseHoliday.Create(
					&entity.EntityHoliday{
						Date: date,
					},
				)

				if err != nil {
					return err
				}
			}

			return nil
		},
	})

	startKafkaConnection(topicParams)
	readTopics()
}
