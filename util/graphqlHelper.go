// Helper Functions for GraphQl data fetching
package util

import (
	"context"
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"net/http"
	"os"
)
// Connects to GraphQL API
func graphQlConnector() (client *graphql.Client,ctx context.Context){
	host := os.Getenv("GRAPH_QL_INTEFRACE")
	httpclient := &http.Client{}
	client = graphql.NewClient(host, graphql.WithHTTPClient(httpclient))
	ctx = context.Background()
	return client,ctx
}
// Bundled Data Fetching
func FetchAllData(repository persistence.Repository) {
	FetchAllModules(repository)
	FetchAllModuleGroups(repository)
	FetchAllDegrees(repository)
	FetchAllProfiles(repository)
}
