package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/usecases"
	"github.com/marcosvieirajr/payment/internal/app/usecases/dto"
)

func CreateTransaction(uc usecases.CreateTransactionUseCase) gin.HandlerFunc {
	var request struct {
		AccountID       int64   `json:"account_id" binding:"required"`
		OperationTypeID int     `json:"operation_type_id" binding:"required"`
		Amount          float64 `json:"amount"`
	}
	return func(c *gin.Context) {
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dto := dto.Transaction{
			Account: dto.Account{
				ID: request.AccountID},
			OperationType: request.OperationTypeID,
			Amount:        request.Amount,
			CreatedFrom:   c.ClientIP()}

		id, err := uc.Execute(c, dto)
		switch err {
		case nil:
		case app.ErrAccountNotFound,
			app.ErrInvalidOperationType,
			app.ErrInvalidAmount:
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
