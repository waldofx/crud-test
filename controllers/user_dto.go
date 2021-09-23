package controllers

import (
	"crud-test/models"
	"time"
)

type UserRequest struct {
	ID		 int	`json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	ClientID int    `json:"client_id"`
}

func (req *UserRequest) toModel() models.User{
	return models.User{
		ID: req.ID,
		Name: req.Name,
		Age: req.Age,
		Sex: req.Sex,
		ClientID: req.ClientID,
	}
}

type UserResponse struct{
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Sex       string `json:"sex"`
	ClientID  int    `json:"client_id"`
}

func newResponse(modelUsers models.User) UserResponse{
	return UserResponse{
		CreatedAt: modelUsers.CreatedAt, 
		UpdatedAt: modelUsers.UpdatedAt,
		DeletedAt: modelUsers.DeletedAt.Time,
		ID: modelUsers.ID,
		Name: modelUsers.Name,
		Age: modelUsers.Age,
		Sex: modelUsers.Sex,
		ClientID: modelUsers.ClientID,
	}
}

func newResponseArray(modelUsers []models.User) []UserResponse{
	var response []UserResponse

	for _, val := range modelUsers{
		response = append(response, newResponse(val))
	}
	return response
}