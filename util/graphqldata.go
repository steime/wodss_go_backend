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

func GetAllModules(repository persistence.Repository){
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(moduleQuery)

	resp := &ModuleResponse{}

	err := client.Run(ctx, req, resp)
	if err != nil {
		log.Fatal(err)
	}

	repository.SaveAllModules(resp.Modules)
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

func GetAllModuleGroups(repository persistence.Repository){
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(moduleGroupQuery)

	resp := &ModuleGroupsResponse{}

	err := client.Run(ctx, req, resp)
	if err != nil {
		log.Fatal(err)
	}

	repository.SaveAllModuleGroups(resp.ModuleGroups)
}
