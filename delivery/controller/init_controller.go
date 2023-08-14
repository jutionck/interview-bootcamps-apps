package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/usecase"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	"net/http"
)

type InitController struct {
	router  *gin.Engine
	useCase usecase.UserUseCase
}

func (i *InitController) initializeHandler(c *gin.Context) {
	userId := common.GenerateID()
	payload := model.User{
		BaseModel: model.BaseModel{Id: userId},
		Username:  "admin",
		Password:  "password",
		Role:      "admin",
		IsActive:  true,
	}

	if err := i.useCase.RegisterNewData(payload); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, nil, "initialized")
}

func NewInitController(r *gin.Engine, uc usecase.UserUseCase) *InitController {
	controller := InitController{
		router:  r,
		useCase: uc,
	}

	rg := r.Group("/api/v1")
	rg.GET("/init", controller.initializeHandler)
	return &controller
}
