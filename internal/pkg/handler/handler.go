package handler

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	router.POST("/update_insurance_amount/", h.issueCalc)

	return router
}
