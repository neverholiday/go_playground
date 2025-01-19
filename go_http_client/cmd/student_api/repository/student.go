package repository

import (
	"context"
	"go_http_client/cmd/student_api/model"

	"github.com/google/uuid"
)

type StudentRepo struct {
	Endpoint   string
	AuthHeader map[string]string
}

func NewStudentRepo(endpoint string, authHeader map[string]string) *StudentRepo {
	return &StudentRepo{
		Endpoint:   endpoint,
		AuthHeader: authHeader,
	}
}

func (r *StudentRepo) CreateStudent(ctx context.Context, req model.StudentCreateRequest) (*model.StudentCreateResponse, error) {
	id := uuid.New().String()
	return &model.StudentCreateResponse{
		ID:   id,
		Name: req.Name,
	}, nil
}
