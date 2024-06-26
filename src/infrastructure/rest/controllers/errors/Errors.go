package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "pausalac/src/domain"
)

type MessagesResponse struct {
	Message string `json:"message"`
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domain.AppError)
		if ok {
			resp := MessagesResponse{Message: err.Error()}
			switch err.Type {
			case domain.NotFound:
				c.JSON(http.StatusNotFound, resp)
				return
			case domain.ValidationError:
				c.JSON(http.StatusBadRequest, resp)
				return
			case domain.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, resp)
				return
			case domain.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, resp)
				return
			case domain.NotAuthorized:
				c.JSON(http.StatusForbidden, resp)
				return
			case domain.RepositoryError:
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
