package util

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
	"net/http"
	"os"
)

/*
var q struct {
	modules struct {
		id graphql.String
	} //`graphql:"module(id: \"9041289\")"`
}

 */


type ModuleResponse struct {
	Modules []persistence.Module `json:"modules"`
}
/*
type Module struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

 */

var moduleQuery = `
query {
  modules {
    id
	name
	code
	credits
	hs
	fs
	requirements {
      id
    }
  }
}`

func GetAllModules(){
	host := os.Getenv("GRAPH_QL_INTEFRACE")
	httpclient := &http.Client{}
	client := graphql.NewClient(host, graphql.WithHTTPClient(httpclient))
	ctx := context.Background()

	req := graphql.NewRequest(moduleQuery)

	resp := &ModuleResponse{}

	err := client.Run(ctx, req, resp)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp.Modules {
		fmt.Println(m)
	}
	/*

	httpclient := &http.Client{}
	client := graphql.NewClient(host,graphql.WithHTTPClient(httpclient))
	ctx := context.Background()
	req := graphql.NewRequest(moduleQuery)

	resp := &ModuleResponse{}

	err := client.Run(ctx, req, resp)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp.Modules {
		fmt.Println(m.ID)
	}

	req := graphql.NewRequest(`
    query ($key: ID!){
		modules  (module:$key){
			module
		}
    }
	`)

	//req.Var("id", "9041289")

	ctx := context.Background()
	var res Resp
	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v",req)
	fmt.Printf("%+v",res)
	spew.Dump(res)


	client := graphql.NewClient(host, nil)

	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)


	query, err := json.Marshal(map[string]string{
		"query":"{modules {\n  id\n }}",
		"variables":"null",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(host,"application/json",bytes.NewBuffer(query))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res))
	modules := &[]persistence.Module{}
	err = UnmarshalGraphQL(res,&result)
	//err = json.Unmarshal(res,modules)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v",modules)
*/
}
