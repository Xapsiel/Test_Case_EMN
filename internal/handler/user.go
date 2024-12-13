package handler

import (
	"github.com/gin-gonic/gin"
	"mobileTest_Case/internal/models"
	"net/http"
)

// @Summary		Регистрация человека
// @Description	Добавление новой песни в базу данных(Обязательные параметры - song,group)
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		models.User	true	"Данные человека"	default({ "nickname": "Worker", "name": "John Doe", "email": "john_doe_1978@gmail.com"})
// @Success		200		{object}	Response
// @Failure		400		{object}	Response
// @Failure		500		{object}	Response
// @Router			/register [post]
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

		new_err := h.service.Delete(input.Email)
		if new_err != nil {
			newErrorResponse(c, 404, new_err.Error())
			return
		}
		newErrorResponse(c, 404, err.Error())
		return
	}

	newResultResponse(c, "Аккаунт зарегистрирован")

}

// @Summary		Подтверждение почты
// @Tags			users
// @Description	Подтверждает почту кодом, присланным на почту
// @Accept			json
// @Produce		json
// @Param			token		query		string	false	"Код подтверждения"
// @Success		200		{object}	Response
// @Failure		400		{object}	Response
// @Failure		500		{object}	Response
// @Router			/verify [get]
func (h *Handler) Verify(c *gin.Context) {
	token := c.Query("token")
	err := h.service.VerifyEmail(token)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newResultResponse(c, "Аккаунт подтвежден")
}
