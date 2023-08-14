package security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jutionck/interview-bootcamp-apps/config"
	"github.com/jutionck/interview-bootcamp-apps/model"
	modelJwt "github.com/jutionck/interview-bootcamp-apps/utils/model"
	"time"
)

type JwtSecurity interface {
	CreateAccessToken(credential model.User) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type jwtSecurity struct {
	cfg config.TokenConfig
}

func (j *jwtSecurity) CreateAccessToken(credential model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(j.cfg.AccessTokenLifeTime)
	claims := &modelJwt.TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: credential.Username,
		Role:     credential.Role,
	}
	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %v", err)
	}
	return ss, nil
}

func (j *jwtSecurity) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != j.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return j.cfg.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != j.cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func NewJwtSecurity(cfg config.TokenConfig) JwtSecurity {
	return &jwtSecurity{cfg: cfg}
}
