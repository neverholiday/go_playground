package main

import (
	"encoding/base64"
	"encoding/json"
	"go_http_client/cmd/student_api/apis"
	"go_http_client/cmd/student_api/model"
	"go_http_client/cmd/student_api/repository"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
)

func main() {

	var envCfg model.EnvCfg
	err := envconfig.Process("GO_HTTP_CLIENT", &envCfg)
	if err != nil {
		panic(err)
	}

	authHeaderData, err := base64.StdEncoding.DecodeString(envCfg.AuthHeader)
	if err != nil {
		panic(err)
	}

	var authHeader map[string]string
	err = json.Unmarshal(authHeaderData, &authHeader)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient

	studentRepo := repository.NewStudentRepo(
		client,
		envCfg.Endpoint,
		authHeader,
	)

	e := echo.New()
	v1g := e.Group("/api/v1")

	apis.
		NewStudentAPI(studentRepo).
		Setup(v1g)

	e.Logger.Fatal(e.Start(":3000"))
}
