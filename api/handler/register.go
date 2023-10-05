package handler

import (
	"fmt"
	"main/api/response"
	"main/models"
	"main/pkg/helper"
	"main/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router       /register [POST]
// @Summary      Create User
// @Description  api for create user
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        data    body     models.CreateUser  true  "data of user"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) Register(c *gin.Context) {

	var req models.CreateUser

	err := c.ShouldBindJSON(&req)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	_, err = h.strg.User().GetByLoging(c.Request.Context(), models.GetByLoginReq{
		Login: req.Login,
	})

	if err == nil {
		fmt.Println("error User already exits", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	hashPass, err := helper.GeneratePasswordHash(req.Password)

	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	req.Password = string(hashPass)

	respCU, err := h.strg.User().Create(c.Request.Context(), req)

	if err != nil {
		fmt.Println("error User Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse{Id: respCU})

}
