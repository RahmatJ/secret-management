package cronConfig

type cronConfig struct {
	DailySecretCheck string
}

var CronConfig = cronConfig{
	// Will check every 00.30 o'clock
	DailySecretCheck: "30 0 * * *",
}
