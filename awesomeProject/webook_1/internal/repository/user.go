package repository

import (
	"awesomeProject/webook_1/internal/dao"
	"awesomeProject/webook_1/internal/domain"
	"context"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
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
		//Nickname: u.Nickname,
		//AboutMe:  u.AboutMe,
		//Birthday: u.Birthday,
	}
}

func (repo UserRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		Id:       u.Id,
		Password: u.Password,
		Nickname: u.Nickname,
		Birthday: u.Birthday,
		AboutMe:  u.AboutMe,
	}
}

//func (repo *UserRepository) UpDomeain(ctx context.Context, u domain.User) error {
//	println(u.Id, "+++++========")
//	return repo.dao.UpdateUser(ctx, dao.User{
//		Id:       u.Id,
//		Birthday: u.Birthday,
//		Nickname: u.Nickname,
//		AboutMe:  u.AboutMe,
//	})
//
//}

func (repo *UserRepository) UpDomeain(ctx context.Context, user domain.User) error {
	err := repo.dao.UpdateUser(ctx, repo.toEntity(user))
	if err != nil {
		return err
	}
	return nil
}
