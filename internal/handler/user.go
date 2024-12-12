package handler

import (
	"github.com/gin-gonic/gin"
	"mobileTest_Case/internal/models"
)

func (h *Handler) Register(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}
	err := h.service.Register(input.Nickname, input.Name, input.Email)
	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}
	go h.service.SendToken(input.Email)
	if err != nil {
		c.AbortWithStatusJSON(404, err)
		return
	}
}
func (h *Handler) Verify(c *gin.Context) {
	token := c.Query("token")
	h.service.VerifyEmail(token)
}
