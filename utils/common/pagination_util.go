package common

import (
	"github.com/jutionck/interview-bootcamp-apps/utils/exception"
	"github.com/jutionck/interview-bootcamp-apps/utils/model"
	"math"
	"os"
	"strconv"
)

func GetPaginationParams(params model.PaginationParam) model.PaginationQuery {
	err := LoadEnv()
	exception.CheckErr(err)

	var (
		page, take, skip int
	)

	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0 {
		n, _ := strconv.Atoi(os.Getenv("DEFAULT_ROWS_PER_PAGE"))
		take = n
	} else {
		take = params.Limit
	}

	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return model.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

func Paginate(page, limit, totalRows int) model.Paging {
	return model.Paging{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}
