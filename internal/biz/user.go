package biz

import (
	"auth-server-boiler-plate/internal/data/models"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepository interface {
	FindById(ctx context.Context, id int32) (*models.User, error)
}

type UserUsecase struct {
	repo UserRepository
	log  *log.Helper
}

func NewUserUsecase(repo UserRepository, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) FindById(ctx context.Context, id int32) (*models.User, error) {
	uc.log.WithContext(ctx).Infof("FindById: %d", id)
	return uc.repo.FindById(ctx, id)
}
