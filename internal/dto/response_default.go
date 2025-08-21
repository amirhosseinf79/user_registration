package dto

import "github.com/gofiber/fiber/v2"

type responseOneMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewDefaultRespose(err error, errStatus int) (int, responseOneMessage) {
	statusCode := fiber.StatusOK
	response := responseOneMessage{
		Code: statusCode,
	}
	if err != nil {
		statusCode = errStatus
		response.Code = statusCode
		response.Message = err.Error()
	}
	return statusCode, response
}
