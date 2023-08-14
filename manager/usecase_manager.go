package manager

import "github.com/jutionck/interview-bootcamp-apps/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	RecruiterUseCase() usecase.RecruiterUseCase
	InterviewerUseCase() usecase.InterviewerUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) RecruiterUseCase() usecase.RecruiterUseCase {
	return usecase.NewRecruiterUseCase(u.repoManager.RecruiterRepo(), u.UserUseCase())
}

func (u *useCaseManager) InterviewerUseCase() usecase.InterviewerUseCase {
	return usecase.NewInterviewerUseCase(u.repoManager.InterviewerRepo(), u.UserUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
