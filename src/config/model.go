package config

type EnvironmentVars struct {
	ISRELEASE bool

	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     int

	KAFKA_BOOTSTRAP_SERVER string
	KAFKA_CLIENT_ID        string
	KAFKA_GROUP_ID         string

	AWS_ACCESS_KEY string
	AWS_SECRET_KEY string

	DEFAULT_ADMIN_MAIL     string
	DEFAULT_ADMIN_PASSWORD string
}
