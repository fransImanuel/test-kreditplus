package handler

import (
	"net/http"
	"test-kreditplus/middleware"
	transaction "test-kreditplus/modules/transaction"
	"test-kreditplus/schemas"
	"test-kreditplus/utils"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionService transaction.Service
}

func InitItemHandler(g *gin.Engine, transactionService transaction.Service) {
	handler := &TransactionHandler{
		TransactionService: transactionService,
	}

	routeAPI := g.Group("/api/v1/transaction")
	routeAPI.Use(middleware.JWTAuthMiddleware())
	{
		routeAPI.POST("/create", handler.CreateTransactionHandler)
	}
}

// Create Transaction
// @Tags Transactions
// @Summary Create Transaction
// @Description Create Transaction
// @ID Transaction-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateTransactionRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/transaction/create [post]
func (h *TransactionHandler) CreateTransactionHandler(c *gin.Context) {
	var req schemas.CreateTransactionRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.TransactionService.CreateTransactionService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create Consumen", map[string]interface{}{
		"id": ID,
	})

}
