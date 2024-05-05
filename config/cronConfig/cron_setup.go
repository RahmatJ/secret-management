package cronConfig

import (
	"fmt"
	cron "github.com/robfig/cron/v3"
	"secret-management/internal/constants"
	"secret-management/internal/domain"
	"time"
)

type CronSetup struct {
	Name          string
	SecretUsecase domain.SecretUsecase
}

func NewCronSetup(secretUsecase domain.SecretUsecase) *CronSetup {
	return &CronSetup{
		Name:          "CronSetup",
		SecretUsecase: secretUsecase,
	}
}

func (cs *CronSetup) Initiate() {
	jakartaTime, _ := time.LoadLocation(constants.Constant.Localization)
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	defer func() {
		fmt.Printf("Stopping Cron")
		scheduler.Stop()
	}()

	fmt.Println("Initiating Cron Job")
	_, err := scheduler.AddFunc(CronConfig.DailySecretCheck, func() {
		err := cs.SecretUsecase.DailySecretCheck()
		if err != nil {
			fmt.Printf("%s: Scheduler failed: %+v", cs.Name, err)
		}
	})
	if err != nil {
		fmt.Printf("%s: Error: %+v", cs.Name, err)
		return
	}

	go scheduler.Start()
}
