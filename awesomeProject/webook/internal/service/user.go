package service

import (
	"awesomeProject/webook/internal/domain"
	"awesomeProject/webook/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateErr          = repository.ErrDuplicateErr
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
	WrongpPssword            = errors.New("密码错误")
	AccountDoesNotExist      = repository.AccountDoesNotExist
)

type UserServer struct {
	repo *repository.UserRepository
}

func NewUserServer(repo *repository.UserRepository) *UserServer {
	return &UserServer{
		repo: repo,
	}
}

func (svc *UserServer) Signup(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return svc.repo.Create(ctx, user)
}

func (svc *UserServer) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//检查密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserServer) Edit(ctx context.Context, user domain.User) error {
	return svc.repo.UpdateUser(ctx, user)
}

func (svc *UserServer) Profile(ctx context.Context, Uid int64) (domain.User, error) {
	u, err := svc.repo.FindUid(ctx, Uid)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (svc *UserServer) LoginJwt(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmailJwt(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, WrongpPssword
	}
	println(u.Id, "==-=-=-=-=-=")
	return u, nil
}
