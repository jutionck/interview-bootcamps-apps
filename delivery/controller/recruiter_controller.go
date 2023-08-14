package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/delivery/middleware"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/model/dto"
	"github.com/jutionck/interview-bootcamp-apps/usecase"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	"github.com/jutionck/interview-bootcamp-apps/utils/exception"
	"net/http"
)

type RecruiterController struct {
	router         *gin.Engine
	uc             usecase.RecruiterUseCase
	authMiddleware middleware.AuthMiddleware
}

func (r *RecruiterController) createHandler(c *gin.Context) {
	var payload dto.RecruiterRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payload.Id = common.GenerateID()
	recruiter := model.Recruiter{
		BaseModel: model.BaseModel{
			Id: payload.Id,
		},
		Name:  payload.Name,
		Email: payload.Email,
		User:  model.User{BaseModel: model.BaseModel{Id: payload.UserId}},
	}

	if err := r.uc.RegisterNewData(recruiter); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreatedResponse(c, payload, "OK")

}
func (r *RecruiterController) listHandler(c *gin.Context) {
	pagingParam, err := exception.ValidateRequestQueryParams(c)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	recruiters, paging, err := r.uc.ListPaging(pagingParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responses []interface{}
	for _, v := range recruiters {
		responses = append(responses, v)
	}
	common.SendPageResponse(c, responses, "OK", paging)
}
func (r *RecruiterController) getHandler(c *gin.Context) {
	id := c.Param("id")
	recruiter, err := r.uc.GetData(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, recruiter, "OK")
}
func (r *RecruiterController) updateHandler(c *gin.Context) {
	var payload dto.RecruiterRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	recruiter := model.Recruiter{
		BaseModel: model.BaseModel{
			Id: payload.Id,
		},
		Name:  payload.Name,
		Email: payload.Email,
		User:  model.User{BaseModel: model.BaseModel{Id: payload.UserId}},
	}

	if err := r.uc.UpdateData(recruiter); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, payload, "OK")
}
func (r *RecruiterController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func NewRecruiterController(
	r *gin.Engine,
	uc usecase.RecruiterUseCase,
	authMiddleware middleware.AuthMiddleware,
) *RecruiterController {
	controller := RecruiterController{
		router:         r,
		uc:             uc,
		authMiddleware: authMiddleware,
	}

	rg := r.Group("/api/v1")
	rg.POST("/recruiters", authMiddleware.RequireToken("admin"), controller.createHandler)
	rg.GET("/recruiters", authMiddleware.RequireToken("admin"), controller.listHandler)
	rg.GET("/recruiters/:id", authMiddleware.RequireToken("admin"), controller.getHandler)
	rg.PUT("/recruiters", authMiddleware.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/recruiters", authMiddleware.RequireToken("admin"), controller.deleteHandler)
	return &controller
}
