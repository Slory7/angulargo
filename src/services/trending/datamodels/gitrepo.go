package datamodels

type GitRepo struct {
	Id            int64
	GitTrendingID int64  `xorm:"'git_trending_id' notnull index"`
	Author        string `xorm:"notnull"`
	Name          string `xorm:"notnull"`
	Href          string `xorm:"notnull"`
	Description   string `xorm:"notnull"`
	Language      string `xorm:"notnull"`
	Stars         int    `xorm:"notnull"`
	Forks         int    `xorm:"notnull"`
}

func (u *GitRepo) TableName() string {
	return "gitrepo"
}
