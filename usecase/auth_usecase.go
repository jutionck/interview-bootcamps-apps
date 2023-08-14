package usecase

import (
	"fmt"
	"github.com/jutionck/interview-bootcamp-apps/utils/security"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
	ActivateAccount(token string) (string, error)
}

type authUseCase struct {
	uc         UserUseCase
	jwtService security.JwtSecurity
}

func (a *authUseCase) Login(username, password string) (string, error) {
	user, err := a.uc.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	if !user.IsActive {
		token, _ := a.uc.GenerateResetToken(user)
		return token, nil
	}

	token, err := a.jwtService.CreateAccessToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authUseCase) ActivateAccount(token string) (string, error) {
	err := a.uc.ActivatedUser(token)
	if err != nil {
		return "", fmt.Errorf("opps, activated user failed")
	}
	return "activated user success", nil
}

func NewAuthUseCase(uc UserUseCase, jwtService security.JwtSecurity) AuthUseCase {
	return &authUseCase{uc: uc, jwtService: jwtService}
}
