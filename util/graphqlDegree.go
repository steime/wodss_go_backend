package util

import (
	"github.com/machinebox/graphql"
	"github.com/steime/wodss_go_backend/persistence"
	"log"
)

type DegreeResponse struct {
	Degrees []persistence.Degree `json:"programs"`
}

var degreeQuery = `
query{
	programs {
		id
		name
		groups {
			id
		}
		profiles {
			id 
		}
	}
}
`

func FetchAllDegrees(repository persistence.Repository) {
	client, ctx := graphQlConnector()

	req := graphql.NewRequest(degreeQuery)

	resp := &DegreeResponse{}

	if err := client.Run(ctx,req,resp); err !=nil {
		log.Print(err)
	} else {
		repository.SaveAllDegrees(resp.Degrees)
	}
}

func UpdateAllDegrees(repository persistence.Repository) func() {
	return func() {
		client, ctx := graphQlConnector()

		req := graphql.NewRequest(degreeQuery)

		resp := &DegreeResponse{}

		if err := client.Run(ctx, req, resp); err != nil {
			log.Print(err)
		} else {
			repository.UpdateAllDegrees(resp.Degrees)
		}
	}
}
