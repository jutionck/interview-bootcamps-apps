package usecase

import (
	"fmt"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/repository"
	"github.com/jutionck/interview-bootcamp-apps/utils/common"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserUseCase interface {
	BaseUseCase[model.User]
	BaseUseCasePaging[model.User]
	FindByUsername(username string) (model.User, error)
	FindByUsernamePassword(username string, password string) (model.User, error)
	GenerateResetToken(payload model.User) (string, error)
	ActivatedUser(token string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNewData(payload model.User) error {
	err := payload.ValidateRequire()
	if err != nil {
		return fmt.Errorf("please input data on required fields : %v", err)
	}
	// username check
	user, _ := u.repo.GetByUsername(payload.Username)
	if user.Username == payload.Username {
		return fmt.Errorf("sorry, username %s already exists", payload.Username)
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	// validation role
	if !payload.IsValidRole() {
		return fmt.Errorf("sorry, user role undefined")
	}
	// save
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	err = u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("oops, create new user fail : %v", err)
	}
	return nil
}

func (u *userUseCase) ListData() ([]model.User, error) {
	return u.repo.List()
}

func (u *userUseCase) GetData(id string) (model.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("sorry, user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) UpdateData(payload model.User) error {
	err := payload.ValidateRequire()
	if err != nil {
		return fmt.Errorf("please input data on required fields : %v", err)
	}
	// username check
	user, _ := u.repo.GetByUsername(payload.Username)
	if user.Username == payload.Username && user.BaseModel.Id != payload.Id {
		return fmt.Errorf("sorry, username %s already exists", payload.Username)
	}
	// password not updated
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	// validation role
	if !payload.IsValidRole() {
		return fmt.Errorf("sorry, user role undefined")
	}
	// save
	payload.UpdatedAt = time.Now()
	err = u.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("oops, update user fail : %v", err)
	}
	return nil
}

func (u *userUseCase) DeleteData(id string) error {
	_, err := u.repo.Get(id)
	if err != nil {
		return fmt.Errorf("sorry, user with ID %s not found", err)
	}
	err = u.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("oops, update user fail : %v", err)
	}
	return nil
}

func (u *userUseCase) ListPaging(requestPaging modelJwt.PaginationParam) ([]model.User, modelJwt.Paging, error) {
	return u.repo.Paging(requestPaging)
}

func (u *userUseCase) FindByUsername(username string) (model.User, error) {
	return u.repo.GetByUsername(username)
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.User, error) {
	return u.repo.GetByUsernamePassword(username, password)
}

func (u *userUseCase) GenerateResetToken(payload model.User) (string, error) {
	token, err := common.GenerateRandomToken(32)
	if err != nil {
		return "", fmt.Errorf("invalid hash token")
	}
	payload.ResetToken = token
	payload.UpdatedAt = time.Now()
	err = u.repo.UpdateResetToken(payload)
	if err != nil {
		return "", fmt.Errorf("oops, update user fail : %v", err)
	}
	return token, nil
}

func (u *userUseCase) ActivatedUser(token string) error {
	return u.repo.ActivatedUser(token)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
