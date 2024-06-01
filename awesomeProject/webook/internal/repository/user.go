package repository

import (
	"awesomeProject/webook/internal/domain"
	"awesomeProject/webook/internal/repository/cache"
	"awesomeProject/webook/internal/repository/dao"
	"context"

	"time"
)

var (
	ErrDuplicateErr     = dao.ErrDuplicateErr
	ErrUserNotFound     = dao.ErrRecocordNotFound
	AccountDoesNotExist = dao.AccountDoesNotExist
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}
func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user domain.User) error {
	err := repo.dao.UpdateUsers(ctx, dao.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Birthday: user.Birthday,
		About_me: user.About_me,
	})
	return err
}

func (repo *UserRepository) FindUid(ctx context.Context, uid int64) (domain.User, error) {
	u, err := repo.cache.Get(ctx, uid)
	u, err = repo.dao.FindByUid(ctx, uid)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toFindByUid(u), nil
}
func (repo *UserRepository) toFindByUid(u dao.User) domain.User {
	return domain.User{
		Email:    u.Email,
		Ctime:    time.Unix(u.Ctime, 0),
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		About_me: u.About_me,
	}
}

func (repo *UserRepository) FindByEmailJwt(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmailjwt(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.todomainJwt(u), nil
}

func (repo *UserRepository) todomainJwt(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}
