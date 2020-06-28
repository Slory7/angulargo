package githubtrending

import (
	"github.com/slory7/angulargo/src/infrastructure/data/repositories"
	m "github.com/slory7/angulargo/src/services/trending/datamodels"
)

type IGithubTrendingService interface {
	New(repo repositories.IRepository, repoReadOnly repositories.IRepositoryReadOnly) IGithubTrendingService
	SaveToDB(trending *m.GitTrendingAll) (exists bool, err error)
}
type GithubTrendingService struct {
	Repository         repositories.IRepository         `inject:"IRepository"`
	RepositoryReadOnly repositories.IRepositoryReadOnly `inject:"IRepositoryReadOnly"`
}

var _ IGithubTrendingService = (*GithubTrendingService)(nil)

func (s *GithubTrendingService) New(repo repositories.IRepository, repoReadOnly repositories.IRepositoryReadOnly) IGithubTrendingService {
	return &GithubTrendingService{repo, repoReadOnly}
}

func (s *GithubTrendingService) SaveToDB(t *m.GitTrendingAll) (exists bool, err error) {
	entity := &m.GitRepoTrending{}
	exists, err = s.RepositoryReadOnly.Exists(entity, "title=?", t.Title)
	if exists {
		return
	}
	dbNew := s.Repository.DB().NewTransaction()
	repoNew := s.Repository.New(dbNew)
	_, err = repoNew.Add(&t.GitRepoTrending)
	if err == nil {
		for _, b := range t.GitRepos {
			b.GitTrendingID = t.GitRepoTrending.Id
			_, err = repoNew.Add(b)
			if err != nil {
				break
			}
		}
		//keep recent only
		if err == nil {
			var nRecent int64 = 30
			if count, _ := repoNew.Count(entity, ""); count > nRecent {
				topSlice := make([]m.GitRepoTrending, 0)
				nTop := int(count - nRecent)
				err = repoNew.DB().ListByCondition(&topSlice, "id", 0, nTop, "", false, nil, "")
				if err == nil {
					topMaxID := topSlice[nTop-1].Id
					_, err = repoNew.DB().DeleteByCondition(entity, "id <= ?", topMaxID)
				}
			}
		}
		if err == nil {
			dbNew.Commit()
		} else {
			dbNew.RollBack()
		}
	}
	return
}
