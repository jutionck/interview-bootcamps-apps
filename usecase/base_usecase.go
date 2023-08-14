package usecase

import "github.com/jutionck/interview-bootcamp-apps/utils/model"

type BaseUseCase[T any] interface {
	RegisterNewData(payload T) error
	ListData() ([]T, error)
	GetData(id string) (T, error)
	UpdateData(payload T) error
	DeleteData(id string) error
}

type BaseUseCasePaging[T any] interface {
	ListPaging(requestPaging model.PaginationParam) ([]T, model.Paging, error)
}
