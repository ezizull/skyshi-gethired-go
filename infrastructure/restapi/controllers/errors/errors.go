// Package errors contains the error handler controller
package errors

import (
	"net/http"

	errorDomains "skyshi_gethired/domain/errors"

	"github.com/gin-gonic/gin"
)

// MessagesResponse is a struct that contains the response body for the message
type MessagesResponse struct {
	Message string `json:"message"`
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*errorDomains.AppError)
		if ok {
			resp := MessagesResponse{Message: err.Error()}
			switch err.Type {
			case errorDomains.NotFound:
				c.JSON(http.StatusNotFound, resp)
				return
			case errorDomains.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case errorDomains.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case errorDomains.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
				return
			case errorDomains.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
				return
			case errorDomains.RepositoryError:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			default:
				c.JSON(http.StatusInternalServerError, MessagesResponse{Message: "We are working to improve the flow of this request."})
				return
			}
		}

		return
	}
}
