package dao //gorm 自带建表功能
import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	//不是优秀的实践，没有走审批流程
	return db.AutoMigrate(&User{})
}
