package util

import (
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
)

type ProfileResponse struct {
	Profiles []persistence.Profile `json:"profiles"`
}

var profileQuery = `
query {
  profiles { 
    id
    name
    min
    modules {
      id
    }
  }
}
`

func FetchAllProfiles(repository persistence.Repository) {
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(profileQuery)

	resp := &ProfileResponse{}

	if err := client.Run(ctx,req,resp); err !=nil {
		log.Print(err)
	} else {
		repository.SaveAllProfiles(resp.Profiles)
	}
}

func UpdateAllProfiles(repository persistence.Repository) func() {
	return func() {
		client, ctx := graphQlConnector()

		req := graphql.NewRequest(profileQuery)

		resp := &ProfileResponse{}

		if err := client.Run(ctx,req,resp); err !=nil {
			log.Print(err)
		} else {
			log.Print("updated")
			repository.UpdateAllProfiles(resp.Profiles)
		}
	}
}
