package handler

import (
	"fmt"
	"main/models"
	"main/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* // @Router       /users [POST]
// @Summary      Create User
// @Description  Create User
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateUser  true  "user data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateUser(c *gin.Context) {

	var user models.CreateUpdateUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.User().Create(c.Request.Context(), user)
	if err != nil {
		fmt.Println("error User Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

} */

// @Router       /users/{id} [put]
// @Summary      Update User
// @Description  api for update users
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user"
// @Param        user    body     models.UpdateUser  true  "data of user"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateUser(c *gin.Context) {

	var user models.UpdateUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	id := c.Param("id")

	resp, err := h.strg.User().Update(c.Request.Context(), id, user)
	if err != nil {
		fmt.Println("error User Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /users/{id} [GET]
// @Summary      Get By Id
// @Description  get user by id
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID" format(uuid)
// @Success      200  {object}  models.User
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetUser(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.strg.User().Get(c.Request.Context(), models.IdRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error User Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /users [GET]
// @Summary      List User
// @Description  get user
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        name    query     string  false  "filter by name"
// @Success      200  {array}   models.User
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllUser(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.strg.User().GetAll(c.Request.Context(), models.GetAllUserRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("name"),
	})
	if err != nil {
		h.log.Error("error User GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /users/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete user by id
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.strg.User().Delete(c.Request.Context(), models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error User Delete:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
