package handler

import (
	"fmt"
	"main/config"
	"main/models"
	"main/pkg/helper"
	"main/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router       /login [GET]
// @Summary      Get Token
// @Description  api for get user token
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        user    body     models.LoginReq  true  "data of user"
// @Success      200  {object}  models.LoginRes
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) Login(c *gin.Context) {

	fmt.Print("regiter is workiing")

	var req models.LoginReq

	err := c.ShouldBindJSON(&req)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	hashPass, err := helper.GeneratePasswordHash(req.Password)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	resp, err := h.strg.User().GetByLoging(c.Request.Context(), models.GetByLoginReq{
		Login: req.Login,
	})

	if err != nil {
		fmt.Println("error Staff GetByLoging:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	err = helper.ComparePasswords([]byte(hashPass), []byte(resp.Password))

	if err != nil {
		h.log.Error("error while compare:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	m := make(map[string]interface{})
	m["user_id"] = resp.Id
	token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, models.LoginRes{Token: token})
}
