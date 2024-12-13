package app

import (
	"context"
	"go_gql/cmd/rickandmorty/model"
	"go_gql/thirdparty/gql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructEpisodeStruct(t *testing.T) {
	expected := []model.Episode{
		{
			ID:   0,
			Name: "revenge of the sift",
		},
		{
			ID:   1,
			Name: "revenge of the surf",
		},
	}

	returnListEpisode := gql.ListEpisodeResponse{
		Episodes: gql.ListEpisodeEpisodes{
			Results: []gql.ListEpisodeEpisodesResultsEpisode{
				{
					Name: "revenge of the sift",
				},
				{
					Name: "revenge of the surf",
				},
			},
		},
	}

	listEpisodeFunc := func(ctx context.Context) (*gql.ListEpisodeResponse, error) {
		return &returnListEpisode, nil
	}

	ctxMock := context.Background()
	builder := NewCsvBuilder(listEpisodeFunc)
	actual, err := builder.ConstructEpisodeStruct(ctxMock)
	assert.Nil(t, err)
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected[0].Name, actual[0].Name)

}
