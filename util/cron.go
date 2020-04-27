// Cron Job for checking GraphQL API
package util

import (
	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/robfig/cron.v2"
)

func CronJob(repository persistence.Repository) {
	c := cron.New()
	c.AddFunc("@daily" , UpdateAllProfiles(repository))
	c.AddFunc("@daily" , UpdateAllModules(repository))
	c.AddFunc("@daily" , UpdateAllDegrees(repository))
	c.AddFunc("@daily" , UpdateAllModuleGroups(repository))
	c.Start()
}
