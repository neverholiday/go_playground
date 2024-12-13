package gql

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

type GqlClient struct {
	Client graphql.Client
}

func NewGqlClient(endpoint string) *GqlClient {
	client := graphql.NewClient(endpoint,
		http.DefaultClient)
	return &GqlClient{Client: client}
}
