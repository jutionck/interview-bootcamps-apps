package manager

import "github.com/jutionck/interview-bootcamp-apps/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	RecruiterRepo() repository.RecruiterRepository
	InterviewerRepo() repository.InterviewerRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) RecruiterRepo() repository.RecruiterRepository {
	return repository.NewRecruiterRepository(r.infra.Conn())
}

func (r *repoManager) InterviewerRepo() repository.InterviewerRepository {
	return repository.NewInterviewerRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
