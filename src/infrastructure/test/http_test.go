package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"

	_ "github.com/crgimenes/goconfig/json"
)

func TestHttpPost(t *testing.T) {
	reqID := uuid.New().String()
	result, err := httpclient.HttpSend("http://api.xxx.cn", "/token", nil, nil, httpclient.FORM, "POST", "", httpclient.TokenEmpty, reqID, true, 5)
	t.Logf("%v", err)
	t.Logf("%v", result)
}
