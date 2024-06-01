package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateErr     = errors.New("邮箱冲突")
	ErrRecocordNotFound = gorm.ErrRecordNotFound //gorm找不到账户的特定错误
	AccountDoesNotExist = gorm.ErrRecordNotFound //gorm特定错误，数据库找不记录
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserdao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Ctime = now
	err := dao.db.WithContext(ctx).Create(&user).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateErr
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	//gorm.ErrRecordNotFound
	var u User
	err := dao.db.WithContext(ctx).Where("Email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) UpdateUsers(ctx context.Context, user User) error {
	return dao.db.WithContext(ctx).Model(&user).Where("id=?", user.Id).Updates(map[string]interface{}{"nickname": user.Nickname, "birthday": user.Birthday, "about_me": user.About_me}).Error
	//	db.Model(&User{}).Where("age < ?", 20).Updates(map[string]interface{}{"name": "Charlie", "age": 18})

}

func (dao *UserDAO) FindByUid(ctx context.Context, uid int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id=?", uid).Find(&u).Error
	return u, err
}

func (dao *UserDAO) FindByEmailjwt(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("Email=?", email).First(&u).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
	Nickname string
	Birthday int
	About_me string
}
