package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"

	_ "github.com/crgimenes/goconfig/json"
)

func TestHttpPost(t *testing.T) {
	app.InitAppInstance(nil)
	reqID := uuid.New().String()
	srv := app.Instance.GetIoCInstanceMust((*httpclient.IHttpClient)(nil)).(httpclient.IHttpClient)
	result, err := srv.HttpSend("http://api.xxx.cn", "/token", nil, nil, httpclient.FORM, "POST", "", httpclient.TokenEmpty, reqID, true, 5)
	t.Logf("%v", err)
	t.Logf("%v", result)
}
