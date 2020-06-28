package datamodels

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type UserWithMemo struct {
	Id          int64
	Memo        null.String
	UpdatedDate time.Time `xorm:"'updated_date' updated"`
}

func (u *UserWithMemo) TableName() string {
	return "users"
}

type UserAll struct {
	User
	UserDetail
	UserLogins []UserLogins
}

type UserWithDetail struct {
	User       `xorm:"extends"`
	UserDetail `xorm:"extends"`
}

func (u *UserWithDetail) TableName() string {
	return "users"
}
