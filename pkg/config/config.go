package config

// "github.com/spf13/viper" // TODO try viper

func GetEnvVar(name string) string {
	switch name {
	case "RMQ_URL":
		// return "amqp://guest:guest@localhost:5001/"
		return "amqp://guest:guest@rabbitmq:5673/"
	case "TEST_QUEUE":
		return "TEST_QUEUE"
	default:
		return ""
	}
}
