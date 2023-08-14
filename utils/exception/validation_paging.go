package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/interview-bootcamp-apps/utils/model"
	"strconv"
)

func ValidateRequestQueryParams(c *gin.Context) (model.PaginationParam, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		return model.PaginationParam{}, fmt.Errorf("invalid page number")
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit <= 0 {
		return model.PaginationParam{}, fmt.Errorf("invalid limit value")
	}

	return model.PaginationParam{
		Page:  page,
		Limit: limit,
	}, nil
}
