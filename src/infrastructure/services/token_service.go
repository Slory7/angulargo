package services

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	"github.com/slory7/angulargo/src/infrastructure/framework/security"
)

type ITokenService interface {
	GetToken() (token httpclient.TokenInfo, err error)
}

type TokenService struct {
	baseURL          string
	tokenRelativeURL string
	clientID         string
	clientSecret     string
	grantType        string
	userName         string
	passWord         string

	cacheKey string
	rwlock   sync.RWMutex
}

var _ ITokenService = (*TokenService)(nil)

func NewTokenService(baseURL string, tokenRelativeURL string, clientID string, clientSecret string, grantType string, userName string, passWord string) *TokenService {
	t := &TokenService{
		baseURL:          baseURL,
		tokenRelativeURL: tokenRelativeURL,
		clientID:         clientID,
		clientSecret:     clientSecret,
		grantType:        grantType,
		userName:         userName,
		passWord:         passWord,
	}
	t.cacheKey = "Token_" + security.ComputeMD5(t.baseURL+"_"+t.tokenRelativeURL+"_"+t.clientID+"_"+t.userName)
	return t
}

func (s *TokenService) GetToken() (token httpclient.TokenInfo, err error) {
	tokenCache := app.Instance.Cache.GetMemoryItem(s.cacheKey, nil, 0)
	if tokenCache != nil {
		return tokenCache.(httpclient.TokenInfo), nil
	}

	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	tokenCache = app.Instance.Cache.GetMemoryItem(s.cacheKey, nil, 0)
	if tokenCache != nil {
		return tokenCache.(httpclient.TokenInfo), nil
	}

	formData := map[string]string{
		"grant_type":    s.grantType,
		"client_id":     s.clientID,
		"client_secret": s.clientSecret,
		"username":      s.userName,
		"password":      s.passWord,
	}
	//srv := app.Instance.GetIoCInstanceMust((*httpclient.IHttpClient)(nil)).(httpclient.IHttpClient)
	srv := app.GetIoCInstanceMust[httpclient.IHttpClient]()
	result, err := srv.HttpPostForm(s.baseURL, s.tokenRelativeURL, formData, httpclient.TokenEmpty, "", true, 5)
	if err == nil && result.IsSuccess {
		var tResult tokenResult
		if err = json.Unmarshal([]byte(result.Content), &tResult); err == nil {
			token.Access_Token = tResult.Access_token
			token.Token_Type = tResult.Token_type
			app.Instance.Cache.SetMemoryItem(s.cacheKey, token, time.Duration(tResult.Expires_in)*time.Second)
		}
	}
	return
}

type tokenResult struct {
	Access_token string
	Token_type   string
	Expires_in   int
}
