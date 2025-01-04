package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func PostUserData(cfg config.Config, email string) error {
	passwordURL := cfg.Auth.BaseURL.JoinPath(cfg.Auth.OtpPath)
	params := passwordURL.Query()
	params.Add("redirect_to", cfg.Auth.OtpRedirectTo)
	passwordURL.RawQuery = params.Encode()

	data := map[string]interface{}{
		"email":       email,
		"data":        nil,
		"create_user": true,
		"gotrue_meta_security": map[string]interface{}{
			"captcha_token": nil,
		},
		"code_challenge":        nil,
		"code_challenge_method": nil,
	}

	_, err := services.PostRequest(
		passwordURL,
		&cfg.Auth.Headers,
		data,
	)
	if err != nil {
		return fmt.Errorf("failed to post user data: %w", err)
	}

	return nil
}

func VerifyUser(cfg config.Config, email string, otp string) (string, error) {
	verifyURL := cfg.Auth.BaseURL.JoinPath(cfg.Auth.VerifyPath)

	data := map[string]interface{}{
		"email":       email,
		"token":       otp,
		"type":        "email",
		"redirect_to": nil,
		"gotrue_meta_security": map[string]interface{}{
			"captchaToken": nil,
		},
	}

	response, err := services.PostRequest(
		verifyURL,
		&cfg.Auth.Headers,
		data,
	)
	if err != nil {
		return "", fmt.Errorf("failed to verify user data: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	accesToken, _ := result["access_token"].(string)

	return accesToken, nil
}

func TradeJWT(cfg config.Config, accessToken string) (string, string, error) {
	tradeURL := cfg.API.BaseURL.JoinPath(cfg.API.TradePath)

	headers := services.MergeHeaders(cfg.API.BasicHeaders, cfg.API.TradeHeaders)
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	UUID, err := services.GenerateUUID()
	if err != nil {
		return "", "", err
	}

	headers.Add("x-device-id", UUID)

	response, err := services.GetRequest(
		tradeURL,
		&headers,
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to trade Supabase access token: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", "", err
	}

	refreshToken, _ := result["jwt"].(string)

	return refreshToken, headers.Get("x-device-id"), nil
}

func RefreshJWT(cfg config.Config) (string, error) {
	refreshURL := cfg.API.BaseURL.JoinPath(cfg.API.RefreshPath)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))
	headers.Add("x-device-id", cfg.Auth.Account.DeviceID)

	response, err := services.GetRequest(
		refreshURL,
		&headers,
	)
	if err != nil {
		return "", fmt.Errorf("failed to refresh your access token: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	accessToken, _ := result["jwt"].(string)

	return accessToken, nil
}
