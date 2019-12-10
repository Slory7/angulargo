package githubtrending

import (
	"github.com/slory7/angulargo/src/services/trending/datamodels"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type IGithubTrendingDocService interface {
	ParseDoc(content string) ([]*datamodels.GitRepo, error)
}

type GithubTrendingDocService struct{}

var _ IGithubTrendingDocService = (*GithubTrendingDocService)(nil)

func (*GithubTrendingDocService) ParseDoc(content string) ([]*datamodels.GitRepo, error) {
	gitRepos := make([]*datamodels.GitRepo, 0)
	reader := strings.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return gitRepos, err
	}

	doc.Find("ol.repo-list li").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h3").Text())
		noSpaceTitle := strings.Replace(title, " ", "", -1)
		author := strings.Split(title, " / ")[0]
		name := strings.Split(title, " / ")[1]
		href := "https://github.com/" + noSpaceTitle
		description := strings.TrimSpace(s.Find(".py-1 p").Text())
		language := strings.TrimSpace(s.Find("[itemprop=programmingLanguage]").Text())
		forkLink := "/" + noSpaceTitle + "/network"
		s1 := s.Find("[href=\"" + forkLink + "\"]").Text()
		s1 = strings.Replace(s1, ",", "", -1)
		s1 = strings.TrimSpace(s1)
		forks, _ := strconv.ParseInt(s1, 10, 32)
		starLink := "/" + noSpaceTitle + "/stargazers"
		s2 := s.Find("[href=\"" + starLink + "\"]").Text()
		s2 = strings.Replace(s2, ",", "", -1)
		s2 = strings.TrimSpace(s2)
		stars, _ := strconv.ParseInt(s2, 10, 32)
		gitRepo := &datamodels.GitRepo{
			Author:      author,
			Name:        name,
			Description: description,
			Href:        href,
			Language:    language,
			Forks:       int(forks),
			Stars:       int(stars),
		}
		gitRepos = append(gitRepos, gitRepo)
	})

	return gitRepos, nil
}
