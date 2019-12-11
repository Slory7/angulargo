package test

import (
	"fmt"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/net/httpclient"
	"github.com/google/uuid"
	"testing"

	_ "github.com/crgimenes/goconfig/json"
)

func TestHttpPost(t *testing.T) {
	reqID := uuid.New().String()
	result, err := httpclient.HttpSend("http://api.xxx.cn", "/token", nil, nil, httpclient.FORM, "POST", "", httpclient.TokenEmpty, reqID, true, 5)
	t.Printf("%v", err)
	t..Printf("%v", result)
}
