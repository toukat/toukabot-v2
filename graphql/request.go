package graphql

import (
	"github.com/toukat/toukabot-v2/config"
	"github.com/toukat/toukabot-v2/util/logger"

	"github.com/shurcooL/graphql"

	"context"
	"fmt"
)

var log = logger.GetLogger("GraphQL")

type ResponseStruct struct {
}

func Query(query string, variables map[string]interface{}, responseStruct *struct{}) (struct{}, error) {
	c := config.GetConfig()
	log.Info("Making GraphQL Query request")

	cl := graphql.NewClient(c.APIHost + "/graphql", nil)

	err := cl.Query(context.Background(), &responseStruct, variables)

	if err != nil {
		log.Error("Unable to make GraphQL request")
		log.Error(fmt.Sprintf("Query: %s", query))
		log.Error(err)
		return *responseStruct, err
	}

	return *responseStruct, nil
}
