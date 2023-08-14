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

type InterviewerController struct {
	router         *gin.Engine
	uc             usecase.InterviewerUseCase
	authMiddleware middleware.AuthMiddleware
}

func (i *InterviewerController) createHandler(c *gin.Context) {
	var payload dto.InterviewerRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payload.Id = common.GenerateID()
	interviewer := model.Interviewer{
		BaseModel: model.BaseModel{
			Id: payload.Id,
		},
		Name:  payload.Name,
		Email: payload.Email,
		User:  model.User{BaseModel: model.BaseModel{Id: payload.UserId}},
	}

	if err := i.uc.RegisterNewData(interviewer); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreatedResponse(c, payload, "OK")

}
func (i *InterviewerController) listHandler(c *gin.Context) {
	pagingParam, err := exception.ValidateRequestQueryParams(c)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	interviewers, paging, err := i.uc.ListPaging(pagingParam)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var responses []interface{}
	for _, v := range interviewers {
		responses = append(responses, v)
	}
	common.SendPageResponse(c, responses, "OK", paging)
}
func (i *InterviewerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	recruiter, err := i.uc.GetData(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(c, recruiter, "OK")
}
func (i *InterviewerController) updateHandler(c *gin.Context) {
	var payload dto.InterviewerRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	interviewer := model.Interviewer{
		BaseModel: model.BaseModel{
			Id: payload.Id,
		},
		Name:  payload.Name,
		Email: payload.Email,
		User:  model.User{BaseModel: model.BaseModel{Id: payload.UserId}},
	}

	if err := i.uc.UpdateData(interviewer); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, payload, "OK")
}
func (i *InterviewerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := i.uc.DeleteData(id); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func NewInterviewerController(
	r *gin.Engine,
	uc usecase.InterviewerUseCase,
	authMiddleware middleware.AuthMiddleware,
) *InterviewerController {
	controller := InterviewerController{
		router:         r,
		uc:             uc,
		authMiddleware: authMiddleware,
	}

	rg := r.Group("/api/v1")
	rg.POST("/interviewers", authMiddleware.RequireToken("admin"), controller.createHandler)
	rg.GET("/interviewers", authMiddleware.RequireToken("admin"), controller.listHandler)
	rg.GET("/interviewers/:id", authMiddleware.RequireToken("admin"), controller.getHandler)
	rg.PUT("/interviewers", authMiddleware.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/interviewers", authMiddleware.RequireToken("admin"), controller.deleteHandler)
	return &controller
}
