package usecase

import (
	"fmt"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/repository"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
)

type RecruiterUseCase interface {
	BaseUseCase[model.Recruiter]
	BaseUseCasePaging[model.Recruiter]
}

type recruiterUseCase struct {
	repo   repository.RecruiterRepository
	userUC UserUseCase
}

func (r *recruiterUseCase) RegisterNewData(payload model.Recruiter) error {
	if err := payload.ValidateRequire(); err != nil {
		return fmt.Errorf("name is required fields")
	}

	recruiter, _ := r.repo.GetByEmail(payload.Email)
	if recruiter.Email == payload.Email {
		return fmt.Errorf("oopss, recruiter with email %v already exists", payload.Email)
	}

	// create user
	payload.User.Username = payload.Email
	payload.User.Role = "recruiter"
	payload.User.Password = "password"
	payload.User.IsActive = false
	err := r.userUC.RegisterNewData(payload.User)
	if err != nil {
		return fmt.Errorf("failed to create user for recruiter %s", payload.Name)
	}

	err = r.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create recruiter")
	}
	return nil
}

func (r *recruiterUseCase) ListData() ([]model.Recruiter, error) {
	return r.repo.List()
}

func (r *recruiterUseCase) GetData(id string) (model.Recruiter, error) {
	return r.repo.Get(id)
}

func (r *recruiterUseCase) UpdateData(payload model.Recruiter) error {
	if err := payload.ValidateRequire(); err != nil {
		return fmt.Errorf("name is required fields")
	}

	recruiter, _ := r.repo.GetByEmail(payload.Email)
	if recruiter.Email == payload.Email && recruiter.Id != payload.Id {
		return fmt.Errorf("oopss, recruiter with email %v already exists", payload.Email)
	}

	// update user
	user, err := r.userUC.FindByUsername(payload.Email)
	if err != nil {
		return fmt.Errorf("user with email %v not found", err)
	}

	user.Username = payload.Email
	err = r.userUC.UpdateData(user)
	if err != nil {
		return fmt.Errorf("failed to update user for recruiter %s", payload.Name)
	}

	err = r.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update recruiter")
	}
	return nil
}

func (r *recruiterUseCase) DeleteData(id string) error {
	_, err := r.repo.Get(id)
	if err != nil {
		return fmt.Errorf("recruiter with ID %v not found", err)
	}

	err = r.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete recruiter")
	}
	return nil
}

func (r *recruiterUseCase) ListPaging(requestPaging modelJwt.PaginationParam) ([]model.Recruiter, modelJwt.Paging, error) {
	return r.repo.Paging(requestPaging)
}

func NewRecruiterUseCase(repo repository.RecruiterRepository, uc UserUseCase) RecruiterUseCase {
	return &recruiterUseCase{repo: repo, userUC: uc}
}
