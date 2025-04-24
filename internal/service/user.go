package service

import (
	v12 "auth-server-boiler-plate/api/v1"
	"auth-server-boiler-plate/internal/biz"
	"context"
)

type UserService struct {
	v12.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}

}

func (s *UserService) GetUser(ctx context.Context, in *v12.GetUserRequest) (*v12.GetUserResponse, error) {
	foundUser, err := s.uc.FindById(ctx, in.Id)

	if err != nil {
		return nil, err
	}

	res := &v12.GetUserResponse{
		Id:       foundUser.ID,
		UserName: foundUser.UserName,
		Email:    foundUser.Email,
		Password: foundUser.Password,
	}

	return res, nil
}
