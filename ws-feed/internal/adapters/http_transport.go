package adapters

import (
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

func NewRestyClient(baseURL string, defaultClientTimeout time.Duration) *resty.Client {
	client := resty.New()

	client.SetBaseURL(baseURL)
	client.SetTimeout(defaultClientTimeout)
	client.SetRetryCount(5)

	json := jsoniter.ConfigFastest
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	return client
}
