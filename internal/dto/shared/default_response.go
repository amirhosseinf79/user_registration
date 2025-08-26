package shared

import "github.com/gofiber/fiber/v2"

type ResponseArgs struct {
	ErrStatus  int
	ErrMessage error
	RealError  error
}

type ResponseOneMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   error  `json:"-"`
}

func NewDefaultResponse(fields ResponseArgs) *ResponseOneMessage {
	statusCode := fiber.StatusOK
	if fields.ErrStatus != 0 {
		statusCode = fields.ErrStatus
	}
	response := ResponseOneMessage{
		Code:  statusCode,
		Error: fields.RealError,
	}
	if fields.ErrMessage != nil {
		response.Message = fields.ErrMessage.Error()
		if fields.RealError == nil {
			response.Error = fields.ErrMessage
		}
	}
	return &response
}
