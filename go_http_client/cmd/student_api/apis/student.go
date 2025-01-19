package apis

import (
	"context"
	"go_http_client/cmd/student_api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IStudentRepo interface {
	CreateStudent(ctx context.Context, req model.StudentCreateRequest, timeout int64) (*model.StudentCreateResponse, error)
}

type StudentAPI struct {
	studentRepo IStudentRepo
}

func NewStudentAPI(studentRepo IStudentRepo) *StudentAPI {
	return &StudentAPI{
		studentRepo: studentRepo,
	}
}

func (s *StudentAPI) Setup(e *echo.Group) {
	e.POST("/students/create", s.createStudent)
}

func (s *StudentAPI) createStudent(c echo.Context) error {

	ctx := c.Request().Context()

	var req model.StudentCreateRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&model.MessageResponse{
				Message: err.Error(),
			},
		)
	}

	resp, err := s.studentRepo.CreateStudent(
		ctx,
		req,
		2, // second
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			&model.MessageResponse{
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, resp)
}
