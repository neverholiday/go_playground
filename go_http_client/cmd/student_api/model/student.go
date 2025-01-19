package model

type StudentCreateRequest struct {
	Name string `json:"name"`
}

type StudentCreateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
