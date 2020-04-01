package util

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
	"os"
)

func graphQlConnector() (client *graphql.Client,ctx context.Context){
	host := os.Getenv("GRAPH_QL_INTEFRACE")
	httpclient := &http.Client{}
	client = graphql.NewClient(host, graphql.WithHTTPClient(httpclient))
	ctx = context.Background()
	return client,ctx
}

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


