package usecase

import (
	"github.com/ismailash/be-enigma-laundry/model/dto"
	"github.com/ismailash/be-enigma-laundry/model/entity"
	"github.com/ismailash/be-enigma-laundry/utils/common"
)

type AuthUseCase interface {
	Register(payload entity.User) (entity.User, error)
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	uc       UserUseCase
	jwtToken common.JwtToken
}

func (a *authUseCase) Register(payload entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.uc.FindByUsernamePassword(payload.Username, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUseCase(uc UserUseCase, jwtToken common.JwtToken) AuthUseCase {
	return &authUseCase{uc: uc, jwtToken: jwtToken}
}
