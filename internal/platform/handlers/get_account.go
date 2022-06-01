package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/marcosvieirajr/payment/internal/app"
	"github.com/marcosvieirajr/payment/internal/app/usecases"
)

func GetAccount(uc usecases.GetAccountUseCase) gin.HandlerFunc {
	var request struct {
		ID int64 `uri:"id"`
	}
	type response struct {
		ID             int64  `json:"account_id"`
		DocumentNumber string `json:"document_number"`
	}

	return func(c *gin.Context) {
		if err := c.BindUri(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(request)
		account, err := uc.Execute(c, request.ID)
		switch err {
		case nil:
		case app.ErrAccountNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var response response
		copier.Copy(&response, account)

		c.JSON(http.StatusOK, response)
	}
}
