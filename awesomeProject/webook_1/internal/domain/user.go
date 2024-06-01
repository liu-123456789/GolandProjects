package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
	Nickname string
	Birthday string
	AboutMe  string
}

//type UserEity struct {
//	Id       int64
//	Nickname string
//	Birthday time.Time
//	AboutMe  string
//}
