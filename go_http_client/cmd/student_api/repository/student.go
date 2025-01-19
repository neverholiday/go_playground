package repository

import (
	"context"
	"go_http_client/cmd/student_api/model"

	"github.com/google/uuid"
)

type StudentRepo struct {
}

func NewStudentRepo() *StudentRepo {
	return &StudentRepo{}
}

func (r *StudentRepo) CreateStudent(ctx context.Context, req model.StudentCreateRequest) (*model.StudentCreateResponse, error) {
	id := uuid.New().String()
	return &model.StudentCreateResponse{
		ID:   id,
		Name: req.Name,
	}, nil
}
