package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/model/dto"
	"github.com/jutionck/interview-bootcamp-apps/usecase"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	"net/http"
)

type AuthController struct {
	router *gin.Engine
	uc     usecase.AuthUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := a.uc.Login(payload.Username, payload.Password)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreatedResponse(c, token, "OK")
}

func (a *AuthController) activatedHandler(c *gin.Context) {
	var resetAccount dto.ResetAccountDto
	if err := c.ShouldBindJSON(&resetAccount); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	rsp, err := a.uc.ActivateAccount(resetAccount.Token)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, rsp, "OK")
}

func NewAuthController(r *gin.Engine, uc usecase.AuthUseCase) *AuthController {
	controller := AuthController{
		router: r,
		uc:     uc,
	}
	rg := r.Group("/api/v1")
	rg.POST("/auth/login", controller.loginHandler)
	rg.POST("/auth/activated", controller.activatedHandler)
	return &controller
}
