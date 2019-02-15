package datamodels

import (
	"time"
)

type GitRepoTrending struct {
	Id           int64
	Title        string    `xorm:"notnull"`
	TrendingDate time.Time `xorm:"'trending_date' notnull"`
}

func (u *GitRepoTrending) TableName() string {
	return "gitrepo_trending"
}
