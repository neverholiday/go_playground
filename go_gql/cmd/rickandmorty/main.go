package main

import (
	"context"
	"go_gql/cmd/rickandmorty/app"
	"go_gql/thirdparty/gql"
)

func main() {
	ctx := context.Background()
	gqlClient := gql.NewGqlClient("https://rickandmortyapi.com/graphql")

	csvBuilder := app.NewCsvBuilder(func(ctx context.Context) (*gql.ListEpisodeResponse, error) {
		return gql.ListEpisode(ctx, gqlClient.Client)
	})
	app := app.NewApp(csvBuilder)

	err := app.Run(ctx)
	if err != nil {
		panic(err)
	}

}
