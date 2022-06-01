package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/usecases"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

func CreateAccount(uc usecases.CreateAccountUseCase) gin.HandlerFunc {
	var request struct {
		DocumentNumber string `json:"document_number" binding:"required"`
	}

	return func(c *gin.Context) {
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dto := dto.Account{
			DocumentNumber: request.DocumentNumber,
			CreatedFrom:    c.ClientIP()}

		id, err := uc.Execute(c, dto)
		switch err {
		case nil:
		case app.ErrDocumentNumberIsInvalid:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		case app.ErrAccountAlreadyExists:
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Writer.Header().Add("Location", fmt.Sprintf("%v/%v", c.Request.RequestURI, *id))
		c.Status(http.StatusCreated)
	}
}
