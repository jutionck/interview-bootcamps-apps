package usecase

import (
	"fmt"
	"github.com/jutionck/interview-bootcamp-apps/model"
	"github.com/jutionck/interview-bootcamp-apps/repository"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
)

type InterviewerUseCase interface {
	BaseUseCase[model.Interviewer]
	BaseUseCasePaging[model.Interviewer]
}

type interviewerUseCase struct {
	repo   repository.InterviewerRepository
	userUC UserUseCase
}

func (i *interviewerUseCase) RegisterNewData(payload model.Interviewer) error {
	if err := payload.ValidateRequire(); err != nil {
		return fmt.Errorf("name is required fields")
	}

	interviewer, _ := i.repo.GetByEmail(payload.Email)
	if interviewer.Email == payload.Email {
		return fmt.Errorf("oopss, interviewer with email %v already exists", payload.Email)
	}

	// create user
	payload.User.Username = payload.Email
	payload.User.Role = "interviewer"
	payload.User.Password = "password"
	payload.User.IsActive = false
	err := i.userUC.RegisterNewData(payload.User)
	if err != nil {
		return fmt.Errorf("failed to create user for interviewer %s", payload.Name)
	}

	err = i.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create interviewer")
	}
	return nil
}

func (i *interviewerUseCase) ListData() ([]model.Interviewer, error) {
	return i.repo.List()
}

func (i *interviewerUseCase) GetData(id string) (model.Interviewer, error) {
	return i.repo.Get(id)
}

func (i *interviewerUseCase) UpdateData(payload model.Interviewer) error {
	if err := payload.ValidateRequire(); err != nil {
		return fmt.Errorf("name is required fields")
	}

	interviewer, _ := i.repo.GetByEmail(payload.Email)
	if interviewer.Email == payload.Email && interviewer.Id != payload.Id {
		return fmt.Errorf("oopss, interviewer with email %v already exists", payload.Email)
	}

	// update user
	user, err := i.userUC.FindByUsername(payload.Email)
	if err != nil {
		return fmt.Errorf("user with email %v not found", err)
	}

	user.Username = payload.Email
	err = i.userUC.UpdateData(user)
	if err != nil {
		return fmt.Errorf("failed to update user for interviewer %s", payload.Name)
	}

	err = i.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update interviewer")
	}
	return nil
}

func (i *interviewerUseCase) DeleteData(id string) error {
	_, err := i.repo.Get(id)
	if err != nil {
		return fmt.Errorf("interviewer with ID %v not found", err)
	}

	err = i.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete interviewer")
	}
	return nil
}

func (i *interviewerUseCase) ListPaging(requestPaging modelJwt.PaginationParam) ([]model.Interviewer, modelJwt.Paging, error) {
	return i.repo.Paging(requestPaging)
}

func NewInterviewerUseCase(repo repository.InterviewerRepository, uc UserUseCase) InterviewerUseCase {
	return &interviewerUseCase{repo: repo, userUC: uc}
}
