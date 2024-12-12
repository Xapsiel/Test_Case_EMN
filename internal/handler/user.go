package handler

import (
	"github.com/gin-gonic/gin"
	"mobileTest_Case/internal/models"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный формат данных")
		return
	}
	err := h.service.Register(input.Nickname, input.Name, input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.service.SendToken(input.Email)
	if err != nil {
		newErrorResponse(c, 404, err.Error())
		return
	}
	newResultResponse(c, "Аккаунт зарегистрирован")

}
func (h *Handler) Verify(c *gin.Context) {
	token := c.Query("token")
	err := h.service.VerifyEmail(token)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResultResponse(c, "Аккаунт подтвежден")
}
