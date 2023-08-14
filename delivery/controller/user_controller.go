package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/delivery/middleware"
	"github.com/jutionck/interview-bootcamp-apps/usecase"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	"github.com/jutionck/interview-bootcamp-apps/utils/exception"
	"net/http"
)

type UserController struct {
	router         *gin.Engine
	uc             usecase.UserUseCase
	authMiddleware middleware.AuthMiddleware
}

func (u *UserController) listHandler(c *gin.Context) {
	paginationParam, err := exception.ValidateRequestQueryParams(c)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	users, paging, err := u.uc.ListPaging(paginationParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range users {
		response = append(response, v)
	}
	common.SendPageResponse(c, response, "OK", paging)
}

func (u *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := u.uc.GetData(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, user, "OK")
}

func NewUserController(r *gin.Engine, uc usecase.UserUseCase, authMiddleware middleware.AuthMiddleware) *UserController {
	controller := UserController{
		router:         r,
		uc:             uc,
		authMiddleware: authMiddleware,
	}

	rg := r.Group("/api/v1")
	rg.GET("/users", authMiddleware.RequireToken("admin"), controller.listHandler)
	rg.GET("/users/:id", authMiddleware.RequireToken("admin"), controller.getHandler)
	return &controller
}
