package util

import (
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
)

type ModuleGroupsResponse struct {
	ModuleGroups []persistence.ModuleGroup `json:"groups"`
}

var moduleGroupQuery = `
query{
  groups {
    id
    name
    minima
    parent{
		id
	}
    modules{
      id
    }
  }
}
`

func FetchAllModuleGroups(repository persistence.Repository){
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(moduleGroupQuery)

	resp := &ModuleGroupsResponse{}

	if err := client.Run(ctx, req, resp);err != nil {
		log.Print(err)
	} else {
		repository.SaveAllModuleGroups(resp.ModuleGroups)
	}
}
