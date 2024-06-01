package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
	Nickname string
	Birthday string
	AboutMe  string
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now

	err := dao.db.WithContext(ctx).Create(&u).Error
	const duplicateErr uint16 = 1062
	if me, ok := err.(*mysql.MySQLError); ok {
		if me.Number == duplicateErr {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {

	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) UpdateUser(ctx context.Context, u User) error {
	fmt.Printf(strconv.FormatInt(u.Id, 10), "=======")
	return dao.db.WithContext(ctx).Model(&u).Where("id=?", u.Id).Updates(map[string]any{
		"utime":    time.Now().UnixMilli(),
		"nickname": u.Nickname,
		"birthday": u.Birthday,
		"about_me": u.AboutMe,
	}).Error
}
