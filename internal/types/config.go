package types

import (
	"net/http"
	"net/url"
)

type API struct {
	BaseURL         *url.URL    `cfg:"config.api.baseUrl,strict"`
	AlbumPath       string      `cfg:"config.api.albumPath,strict"`
	SeriesPath      string      `cfg:"config.api.seriesPath,strict"`
	AccountPath     string      `cfg:"config.api.accountPath,strict"`
	ProfilePath     string      `cfg:"config.api.profilePath,strict"`
	TradePath       string      `cfg:"config.api.tradePath,strict"`
	RefreshPath     string      `cfg:"config.api.refreshPath,strict"`
	BasicHeaders    http.Header `cfg:"config.api.basicHeaders,strict"`
	TradeHeaders    http.Header `cfg:"config.api.tradeHeaders,strict"`
	PlaybookHeaders http.Header `cfg:"config.api.playbookHeaders,strict"`
}

type Auth struct {
	BaseURL       *url.URL    `cfg:"config.auth.baseUrl,strict"`
	VerifyPath    string      `cfg:"config.auth.verifyPath,strict"`
	OtpPath       string      `cfg:"config.auth.otpPath,strict"`
	OtpRedirectTo string      `cfg:"config.auth.otpRedirectTo,strict"`
	Headers       http.Header `cfg:"config.auth.headers,strict"`
	Account       AuthAccount
}

type AuthAccount struct {
	AccessToken  string `cfg:"config.auth.account.accessToken,strict"`
	RefreshToken string `cfg:"config.auth.account.refreshToken,strict"`
	DeviceID     string `cfg:"config.auth.account.deviceId,strict"`
}
