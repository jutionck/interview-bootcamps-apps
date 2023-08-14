package repository

import (
	"database/sql"
	"fmt"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	modelUtils "github.com/jutionck/interview-bootcamp-apps/utils/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	BaseRepository[model.User]
	BaseRepositoryPaging[model.User]
	GetByUsername(username string) (model.User, error)
	GetByUsernamePassword(username string, password string) (model.User, error)
	UpdateResetToken(payload model.User) error
	ActivatedUser(token string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.User) error {
	_, err := u.db.Exec(
		common.UserCreate,
		payload.Id,
		payload.Username,
		payload.Password,
		payload.IsActive,
		payload.Role,
		payload.CreatedAt,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Paging(requestPaging modelUtils.PaginationParam) ([]model.User, modelUtils.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)
	rows, err := u.db.Query(common.UserListPaging,
		paginationQuery.Take,
		paginationQuery.Skip,
	)
	if err != nil {
		return nil, modelUtils.Paging{}, err
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.BaseModel.Id,
			&user.Username,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, modelUtils.Paging{}, err
		}
		users = append(users, user)
	}
	var totalRows int
	row := u.db.QueryRow("SELECT COUNT(*) FROM users")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, modelUtils.Paging{}, err
	}
	paging := common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows)
	return users, paging, nil
}

func (u *userRepository) List() ([]model.User, error) {
	rows, err := u.db.Query(common.UserList)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.BaseModel.Id,
			&user.Username,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.UserGet, id).Scan(
		&user.BaseModel.Id,
		&user.Username,
		&user.IsActive,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) Update(payload model.User) error {
	_, err := u.db.Exec(
		common.UserUpdate,
		payload.BaseModel.Id,
		payload.Username,
		payload.Password,
		payload.IsActive,
		payload.Role,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id string) error {
	_, err := u.db.Exec(common.UserDelete, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(
		common.UserGetUsernamePassword,
		username,
	).Scan(
		&user.BaseModel.Id,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.IsActive,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetByUsernamePassword(username string, password string) (model.User, error) {
	user, err := u.GetByUsername(username)
	if err != nil {
		return model.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("failed to verify password hash : %v", err)
	}
	return user, nil
}

func (u *userRepository) UpdateResetToken(payload model.User) error {
	_, err := u.db.Exec(
		common.UserResetToken,
		payload.Id,
		payload.ResetToken,
		payload.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) ActivatedUser(token string) error {
	_, err := u.db.Exec(
		common.UserActivate,
		token,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
