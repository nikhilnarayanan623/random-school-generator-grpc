package handler

import (
	"api-gateway/pkg/api/handler/interfaces"
	clientInterface "api-gateway/pkg/client/interfaces"
)

type schoolHandler struct {
	client clientInterface.SchoolClient
}

func NewSchoolHandler(client clientInterface.SchoolClient) interfaces.SchoolHandler {
	return &schoolHandler{
		client: client,
	}
}
