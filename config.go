package tiktokbiz

import "net/http"

const (
	BaseAPIURL        = "https://business-api.tiktok.com/open_api"
	BaseSandboxAPIURL = "https://sandbox-ads.tiktok.com/open_api"
)

type APIVersion string

const (
	APIVersion12 APIVersion = "v1.2"
	APIVersion13 APIVersion = "v1.3"
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	appID      string
	secret     string
	AuthToken  string
	BaseURL    string
	APIVersion APIVersion
	HTTPClient *http.Client
}

func DefaultConfig(appID, secret string) ClientConfig {
	return ClientConfig{
		appID:      appID,
		secret:     secret,
		BaseURL:    BaseAPIURL,
		APIVersion: APIVersion13,
		HTTPClient: &http.Client{},
	}
}

func (ClientConfig) String() string {
	return "<Tiktok Business API ClientConfig>"
}
