package test

import (
	"fmt"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/net/httpclient"
	"testing"

	_ "github.com/crgimenes/goconfig/json"
	"github.com/satori/go.uuid"
)

func TestHttpPost(t *testing.T) {
	reqID := uuid.NewV4().String()
	result, err := httpclient.HttpSend("http://api.xxx.cn", "/token", nil, nil, httpclient.FORM, "POST", "", httpclient.TokenEmpty, reqID, true, 5)
	t.Printf("%v", err)
	t..Printf("%v", result)
}
