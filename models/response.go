package models

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int                    `json:"-"`
	Success    bool                   `json:"success"`
	Message    string                 `json:"message,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
}

func (response *Response) SendResponse(c *gin.Context) {
	c.AbortWithStatusJSON(response.StatusCode, response)
}

func SendResponseData(c *gin.Context, statusCode int, data map[string]interface{}) {
	response := &Response{
		StatusCode: statusCode,
		Success:    true,
		Data:       data,
	}
	response.SendResponse(c)
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	response := &Response{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
	}
	response.SendResponse(c)
}
