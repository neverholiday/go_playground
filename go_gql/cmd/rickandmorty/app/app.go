package app

import (
	"context"
	"fmt"
	"go_gql/cmd/rickandmorty/model"
	"go_gql/thirdparty/gql"

	"github.com/gocarina/gocsv"
)

type ListEpisodeMethod func(ctx context.Context) (*gql.ListEpisodeResponse, error)

type CSvBuilder struct {
	ListEpisodeFunc ListEpisodeMethod
}

func NewCsvBuilder(listEpisodeFunc ListEpisodeMethod) *CSvBuilder {
	return &CSvBuilder{ListEpisodeFunc: listEpisodeFunc}
}

func (b *CSvBuilder) ConstructEpisodeStruct(ctx context.Context) ([]model.Episode, error) {
	response, err := b.ListEpisodeFunc(ctx)
	if err != nil {
		return nil, fmt.Errorf("err : %v", err)
	}

	var episodes []model.Episode
	for i, char := range response.Episodes.Results {
		episodes = append(episodes, model.Episode{
			ID:   int64(i),
			Name: char.GetName(),
		})
	}
	return episodes, nil
}

type App struct {
	csvBuilder *CSvBuilder
}

func NewApp(b *CSvBuilder) *App {
	return &App{csvBuilder: b}
}

func (a *App) generateCSV(ctx context.Context) (string, error) {
	episodes, err := a.csvBuilder.ConstructEpisodeStruct(ctx)
	if err != nil {
		return "", fmt.Errorf("err: %v", err)
	}

	epCsv, err := gocsv.MarshalString(episodes)
	if err != nil {
		return "", fmt.Errorf("err: %v", err)
	}

	return epCsv, nil
}

func (a *App) Run(ctx context.Context) error {
	csv, err := a.generateCSV(ctx)
	if err != nil {
		return fmt.Errorf("err: %v", err)
	}

	fmt.Println(csv)
	return nil
}
