package repository

import (
	"database/sql"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
)

type RecruiterRepository interface {
	BaseRepository[model.Recruiter]
	BaseRepositoryPaging[model.Recruiter]
	GetByEmail(email string) (model.Recruiter, error)
}

type recruiterRepository struct {
	db *sql.DB
}

func (r *recruiterRepository) Create(payload model.Recruiter) error {
	_, err := r.db.Exec(
		`INSERT INTO recruiter (id, name, email, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
		payload.Id,
		payload.Name,
		payload.Email,
		payload.User.Id,
		payload.CreatedAt,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *recruiterRepository) List() ([]model.Recruiter, error) {
	rows, err := r.db.Query(`SELECT id, name, email, created_at, updated_at FROM recruiter ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	var recruiters []model.Recruiter
	for rows.Next() {
		var recruiter model.Recruiter
		err := rows.Scan(
			&recruiter.Id,
			&recruiter.Name,
			&recruiter.Email,
			&recruiter.CreatedAt,
			&recruiter.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		recruiters = append(recruiters, recruiter)
	}
	return recruiters, nil
}

func (r *recruiterRepository) Get(id string) (model.Recruiter, error) {
	var recruiter model.Recruiter
	err := r.db.QueryRow(`SELECT id, name, email, created_at, updated_at FROM recruiter WHERE id = $1`, id).Scan(
		&recruiter.Id,
		&recruiter.Name,
		&recruiter.Email,
		&recruiter.CreatedAt,
		&recruiter.UpdatedAt,
	)
	if err != nil {
		return model.Recruiter{}, err
	}
	return recruiter, nil
}

func (r *recruiterRepository) Update(payload model.Recruiter) error {
	_, err := r.db.Exec(
		`UPDATE recruiter SET name = $2, email = $3, updated_at = $4 WHERE id = $1`,
		payload.Id,
		payload.Name,
		payload.Email,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *recruiterRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM recruiter WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *recruiterRepository) Paging(requestPaging modelJwt.PaginationParam) ([]model.Recruiter, modelJwt.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := r.db.Query(`SELECT id, name, email, created_at, updated_at FROM recruiter ORDER BY created_at DESC LIMIT $1 OFFSET $2`, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, modelJwt.Paging{}, err
	}
	var recruiters []model.Recruiter
	for rows.Next() {
		var recruiter model.Recruiter
		err := rows.Scan(&recruiter.Id, &recruiter.Name, &recruiter.Email, &recruiter.CreatedAt, &recruiter.UpdatedAt)
		if err != nil {
			return nil, modelJwt.Paging{}, err
		}
		recruiters = append(recruiters, recruiter)
	}
	var totalRows int
	err = r.db.QueryRow("SELECT COUNT(*) FROM recruiter").Scan(&totalRows)
	if err != nil {
		return nil, modelJwt.Paging{}, err
	}
	return recruiters, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (r *recruiterRepository) GetByEmail(email string) (model.Recruiter, error) {
	var recruiter model.Recruiter
	err := r.db.QueryRow(`SELECT id, name, email, created_at, updated_at FROM recruiter WHERE email = $1`, email).Scan(
		&recruiter.Id,
		&recruiter.Name,
		&recruiter.Email,
		&recruiter.CreatedAt,
		&recruiter.UpdatedAt,
	)
	if err != nil {
		return model.Recruiter{}, err
	}
	return recruiter, nil
}

func NewRecruiterRepository(db *sql.DB) RecruiterRepository {
	return &recruiterRepository{db: db}
}
