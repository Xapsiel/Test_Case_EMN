package handler

import (
	"github.com/gin-gonic/gin"
	"mobileTest_Case/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/register", h.Register)
	router.GET("/verify", h.Verify)
	return router
}
