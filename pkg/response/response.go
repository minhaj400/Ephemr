package response

import (
	"log"
	"net/http"

	"github.com/Minhajxdd/Ephemr/pkg/errs"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message,omitempty"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
	Meta    any        `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Details string `json:"details"`
}

func Success(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func List(ctx *gin.Context, data any, meta any) {
	ctx.JSON(http.StatusOK, BaseResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func HandleError(ctx *gin.Context, err error) {
	if appErr, ok := errs.From(err); ok {
		if appErr.Internal != nil {
			//here implement robust error handler
			log.Printf("[ERROR] %s: %v", appErr.Code, appErr.Internal)
		}
		ctx.JSON(appErr.HTTPStatus, BaseResponse{
			Success: false,
			Error: &ErrorInfo{
				Code:    appErr.Code,
				Details: appErr.Message,
			},
		})
		return
	}

	log.Printf("[ERROR] %v", err)
	ctx.JSON(http.StatusInternalServerError, BaseResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    "INTERNAL_SERVER_ERROR",
			Details: "Something went wrong",
		},
	})
}
