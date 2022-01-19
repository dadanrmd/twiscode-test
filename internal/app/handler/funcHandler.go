package handler

import (
	"twiscode-test/internal/app/commons/jsonHttpResponse"

	"github.com/gin-gonic/gin"
)

type FuncHandler struct {
	HandlerOption
}

func (handler FuncHandler) ConvertBinaryToDecimal(c *gin.Context) {
	inp := c.Query("key")
	if inp == "" {
		errPayload := jsonHttpResponse.FailedResponse{
			Status:       jsonHttpResponse.FailedStatus,
			ResponseCode: "00",
			Message:      "params empty",
		}
		jsonHttpResponse.BadRequest(c, errPayload)
		return
	}
	res := handler.FuncService.BinaryToDecimal(inp)
	successPayload := jsonHttpResponse.SuccessResponse{
		Status:       jsonHttpResponse.SuccessStatus,
		ResponseCode: "200",
		Message:      "",
		Data:         res,
	}
	jsonHttpResponse.OK(c, successPayload)
}

func (handler FuncHandler) ConvertDecimalToBinary(c *gin.Context) {
	inp := c.Query("key")
	if inp == "" {
		errPayload := jsonHttpResponse.FailedResponse{
			Status:       jsonHttpResponse.FailedStatus,
			ResponseCode: "00",
			Message:      "params empty",
		}
		jsonHttpResponse.BadRequest(c, errPayload)
		return
	}
	res := handler.FuncService.DecimalToBinary(inp)

	successPayload := jsonHttpResponse.SuccessResponse{
		Status:       jsonHttpResponse.SuccessStatus,
		ResponseCode: "200",
		Message:      "",
		Data:         res,
	}
	jsonHttpResponse.OK(c, successPayload)
}

func (handler FuncHandler) GetPolyndrome(c *gin.Context) {
	inp := c.Query("key")
	if inp == "" {
		errPayload := jsonHttpResponse.FailedResponse{
			Status:       jsonHttpResponse.FailedStatus,
			ResponseCode: "00",
			Message:      "params empty",
		}
		jsonHttpResponse.BadRequest(c, errPayload)
		return
	}
	res := handler.FuncService.Polyndrome(inp)
	successPayload := jsonHttpResponse.SuccessResponse{
		Status:       jsonHttpResponse.SuccessStatus,
		ResponseCode: "200",
		Message:      "",
		Data:         res,
	}
	jsonHttpResponse.OK(c, successPayload)
}
