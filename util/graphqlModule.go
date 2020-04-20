package util

import (
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
)

type ModuleResponse struct {
	Modules []persistence.Module `json:"modules"`
}

var moduleQuery = `
query {
  modules {
    id
	name
	code
	credits
	hs
	fs
	msp
	requirements {
      id
    }
  }
}`

func FetchAllModules(repository persistence.Repository){
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(moduleQuery)

	resp := &ModuleResponse{}

	if err := client.Run(ctx, req, resp); err != nil {
		log.Print(err)
	} else {
		repository.SaveAllModules(resp.Modules)
	}
}
