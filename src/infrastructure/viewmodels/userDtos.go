package viewmodels

import "time"
import "gopkg.in/guregu/null.v3"

type UserWithID struct {
	Id int64 `validate:"gt=0"`
}

type UserInfoDto struct {
	UserName string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      uint8  `validate:"gte=0,lte=130"`
	Memo     null.String
}

type UserDto struct {
	UserWithID
	UserInfoDto
}

type UserMemoDto struct {
	UserWithID
	Memo null.String
}

type UserAllDto struct {
	UserDto
	UserDetailDto
	UserLogins []UserLoginDto
}

type UserWithDetailDto struct {
	UserDto
	UserDetailDto
}

type UserDetailDto struct {
	LastLoginTime time.Time
}

type UserLoginDto struct {
	LoginTime time.Time
}
