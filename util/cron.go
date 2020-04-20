package util

import (
	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/robfig/cron.v2"
)

func CronJob(repository persistence.Repository) {
	c := cron.New()
	c.AddFunc("@every 1m" , UpdateAllProfiles(repository))
	c.AddFunc("@every 1m" , UpdateAllModules(repository))
	c.AddFunc("@every 1m" , UpdateAllDegrees(repository))
	c.AddFunc("@every 1m" , UpdateAllModuleGroups(repository))
	c.Start()
}
