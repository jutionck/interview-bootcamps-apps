package repository

import (
	"database/sql"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
)

type InterviewerRepository interface {
	BaseRepository[model.Interviewer]
	BaseRepositoryPaging[model.Interviewer]
	GetByEmail(email string) (model.Interviewer, error)
}

type interviewerRepository struct {
	db *sql.DB
}

func (i *interviewerRepository) Create(payload model.Interviewer) error {
	_, err := i.db.Exec(
		`INSERT INTO interviewer (id, name, email, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`,
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

func (i *interviewerRepository) List() ([]model.Interviewer, error) {
	rows, err := i.db.Query(`SELECT id, name, email, created_at, updated_at FROM interviewer ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	var interviewers []model.Interviewer
	for rows.Next() {
		var interviewer model.Interviewer
		err := rows.Scan(
			&interviewer.Id,
			&interviewer.Name,
			&interviewer.Email,
			&interviewer.CreatedAt,
			&interviewer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		interviewers = append(interviewers, interviewer)
	}
	return interviewers, nil
}

func (i *interviewerRepository) Get(id string) (model.Interviewer, error) {
	var interviewer model.Interviewer
	err := i.db.QueryRow(`SELECT id, name, email, created_at, updated_at FROM interviewer WHERE id = $1`, id).Scan(
		&interviewer.Id,
		&interviewer.Name,
		&interviewer.Email,
		&interviewer.CreatedAt,
		&interviewer.UpdatedAt,
	)
	if err != nil {
		return model.Interviewer{}, err
	}
	return interviewer, nil
}

func (i *interviewerRepository) Update(payload model.Interviewer) error {
	_, err := i.db.Exec(
		`UPDATE interviewer SET name = $2, email = $3, updated_at = $4 WHERE id = $1`,
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

func (i *interviewerRepository) Delete(id string) error {
	_, err := i.db.Exec(`DELETE FROM interviewer WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (i *interviewerRepository) Paging(requestPaging modelJwt.PaginationParam) ([]model.Interviewer, modelJwt.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := i.db.Query(`SELECT id, name, email, created_at, updated_at FROM interviewer ORDER BY created_at DESC LIMIT $1 OFFSET $2`, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, modelJwt.Paging{}, err
	}
	var interviewers []model.Interviewer
	for rows.Next() {
		var interviewer model.Interviewer
		err := rows.Scan(
			&interviewer.Id,
			&interviewer.Name,
			&interviewer.Email,
			&interviewer.CreatedAt,
			&interviewer.UpdatedAt,
		)
		if err != nil {
			return nil, modelJwt.Paging{}, err
		}
		interviewers = append(interviewers, interviewer)
	}
	var totalRows int
	err = i.db.QueryRow("SELECT COUNT(*) FROM interviewer").Scan(&totalRows)
	if err != nil {
		return nil, modelJwt.Paging{}, err
	}
	return interviewers, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (r *interviewerRepository) GetByEmail(email string) (model.Interviewer, error) {
	var interviewer model.Interviewer
	err := r.db.QueryRow(`SELECT id, name, email, created_at, updated_at FROM interviewer WHERE email = $1`, email).Scan(
		&interviewer.Id,
		&interviewer.Name,
		&interviewer.Email,
		&interviewer.CreatedAt,
		&interviewer.UpdatedAt,
	)
	if err != nil {
		return model.Interviewer{}, err
	}
	return interviewer, nil
}

func NewInterviewerRepository(db *sql.DB) InterviewerRepository {
	return &interviewerRepository{db: db}
}
