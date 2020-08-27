package app

import (
	"github.com/gin-gonic/gin"
	"github.com/goudai-projects/gd-ops/model"
	"net/http"
)

func JSONAppErrorHandler() gin.HandlerFunc {
	return jsonAppErrorHandlerT(gin.ErrorTypeAny)
}

func jsonAppErrorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *model.AppError
			switch err.(type) {
			case *model.AppError:
				parsedError = err.(*model.AppError)
			default:
				parsedError = model.NewAppError("app.error.unknown", nil, http.StatusInternalServerError)
			}
			// Put the error into response
			lang := c.Param("lang")
			accept := c.GetHeader("Accept-Language")
			parsedError.Localization(lang, accept)
			c.AbortWithStatusJSON(parsedError.StatusCode, parsedError)
			return
		}

	}
}
