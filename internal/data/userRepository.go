package data

import (
	"auth-server-boiler-plate/internal/biz"
	"auth-server-boiler-plate/internal/data/models"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepository struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepository(data *Data, logger log.Logger) biz.UserRepository {
	return &userRepository{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepository) FindById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := r.data.db.WithContext(ctx).First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
