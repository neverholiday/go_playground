package main

import (
	"go_http_client/cmd/student_api/apis"
	"go_http_client/cmd/student_api/repository"

	"github.com/labstack/echo/v4"
)

func main() {

	studentRepo := repository.NewStudentRepo()

	e := echo.New()
	g := e.Group("/api/v1")

	apis.
		NewStudentAPI(studentRepo).
		Setup(g)

	e.Logger.Fatal(e.Start(":3000"))
}
