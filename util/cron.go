// Cron Job for checking GraphQL API
package util

import (
	"github.com/steime/wodss_go_backend/persistence"
	"gopkg.in/robfig/cron.v2"
	"log"
)

func CronJob(repository persistence.Repository) {
	c := cron.New()
	if _, err := c.AddFunc("@daily" , UpdateAllProfiles(repository)); err != nil {
		log.Print(err)
	}
	if _, err := c.AddFunc("@daily" , UpdateAllModules(repository)); err != nil {
		log.Print(err)
	}
	if _, err := c.AddFunc("@daily" , UpdateAllDegrees(repository)); err != nil {
		log.Print(err)
	}
	if _, err := c.AddFunc("@daily" , UpdateAllModuleGroups(repository)); err != nil {
		log.Print(err)
	}
	c.Start()
}
