package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go_http_client/cmd/student_api/model"
	"io"
	"net/http"
	"time"
)

type StudentRepo struct {
	Client     *http.Client
	Endpoint   string
	AuthHeader map[string]string
}

func NewStudentRepo(client *http.Client, endpoint string, authHeader map[string]string) *StudentRepo {
	return &StudentRepo{
		Client:     client,
		Endpoint:   endpoint,
		AuthHeader: authHeader,
	}
}

func (r *StudentRepo) CreateStudent(ctx context.Context, req model.StudentCreateRequest, timeout int64) (*model.StudentCreateResponse, error) {

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(
			timeout*int64(time.Second),
		),
	)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		r.Endpoint,
		bytes.NewBuffer(reqJson),
	)
	if err != nil {
		return nil, err
	}

	for k, v := range r.AuthHeader {
		httpReq.Header.Add(k, v)
	}

	httpResp, err := r.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != 200 {

		var messageResp struct {
			Detail string `json:"detail"`
		}

		err = json.Unmarshal(body, &messageResp)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(messageResp.Detail)
	}

	var resp model.StudentCreateResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
