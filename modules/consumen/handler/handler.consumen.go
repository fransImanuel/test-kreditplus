package handler

import (
	"net/http"
	"test-kreditplus/middleware"
	consumen "test-kreditplus/modules/consumen"
	"test-kreditplus/schemas"
	"test-kreditplus/utils"

	"github.com/gin-gonic/gin"
)

type ConsumenHandler struct {
	ConsumenService consumen.Service
}

func InitConsumenHandler(g *gin.Engine, consumenService consumen.Service) {
	handler := &ConsumenHandler{
		ConsumenService: consumenService,
	}

	routeAPI := g.Group("/api/v1/consumen")
	routeAPI.Use(middleware.JWTAuthMiddleware())
	{

		routeAPI.GET("/", handler.GetAllConsumenHandler)
		routeAPI.POST("/create", handler.CreateConsumenHandler)
	}
}

// Create Consumen
// @Tags Consumens
// @Summary Create Consumen
// @Description Create Consumen
// @ID Consumen-Create
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.CreateConsumenRequest true "body data"
// @Success 200  {object} schemas.Response
// @Router /v1/consumen/create [post]
func (h *ConsumenHandler) CreateConsumenHandler(c *gin.Context) {
	var req schemas.CreateConsumenRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, "Bad Request", "Required field is empty", nil)
		return
	}

	err, ID := h.ConsumenService.CreateConsumenService(req)
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Create Consumen", map[string]interface{}{
		"id": ID,
	})

}

// Get Consumen
// @Tags Consumens
// @Summary Get Consumen
// @Description Get Consumen
// @ID Consumen-Get
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200  {object} schemas.Response
// @Router /v1/consumen/ [get]
func (h *ConsumenHandler) GetAllConsumenHandler(c *gin.Context) {
	consumens, err := h.ConsumenService.GetConsumenService()
	if err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, "Error", err.Error(), nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, "success", "Success Get Consumens", consumens)
}
