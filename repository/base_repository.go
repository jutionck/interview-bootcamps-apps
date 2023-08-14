package repository

import "github.com/jutionck/interview-bootcamp-apps/utils/model"

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
	Get(id string) (T, error)
	Update(payload T) error
	Delete(id string) error
}

type BaseRepositoryPaging[T any] interface {
	Paging(requestPaging model.PaginationParam) ([]T, model.Paging, error)
}
