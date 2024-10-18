package initialize


func Run() {
	// 1. Load config
	LoadConfig()
	// 2. Init logger
	InitLogger()
	// 3. Config Database
	InitDatabase()
	// 4. Config Kafka (or RabbitMQ)

	// 5. Config Redis Cache

	// 6. Init API


}