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

	doc.Find("article.Box-row").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h1").Text())
		noSpaceTitle := strings.Replace(title, "\n", "", -1)
		noSpaceTitle = strings.Replace(noSpaceTitle, " ", "", -1)
		author := strings.Split(noSpaceTitle, "/")[0]
		name := strings.Split(noSpaceTitle, "/")[1]
		href := "https://github.com/" + noSpaceTitle
		description := strings.TrimSpace(s.Find("p").Text())
		language := strings.TrimSpace(s.Find("[itemprop=programmingLanguage]").Text())

		starHref := "/" + noSpaceTitle + "/stargazers." + name
		sstar := strings.TrimSpace(s.Find("[href=\"" + starHref + "\"]").Text())
		stars, _ := strconv.ParseInt(strings.Replace(sstar, ",", "", -1), 10, 32)

		forkHref := "/" + noSpaceTitle + "/network/members." + name
		sfork := strings.TrimSpace(s.Find("[href=\"" + forkHref + "\"]").Text())
		forks, _ := strconv.ParseInt(strings.Replace(sfork, ",", "", -1), 10, 32)

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
