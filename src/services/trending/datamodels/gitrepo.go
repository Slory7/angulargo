package datamodels

type GitRepo struct {
	Id            int64
	GitTrendingID int64  `xorm:"'git_trending_id' notnull index"`
	Author        string `xorm:"notnull" ,validate:"required"`
	Name          string `xorm:"notnull" ,validate:"required"`
	Href          string `xorm:"notnull" ,validate:"required"`
	Description   string `xorm:"varchar(2000) notnull" ,validate:"required"`
	Language      string `xorm:"notnull" ,validate:"required"`
	Stars         int    `xorm:"notnull"`
	Forks         int    `xorm:"notnull"`
}

func (u *GitRepo) TableName() string {
	return "gitrepo"
}
