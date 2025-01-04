package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func GetAccount(cfg config.Config, accessToken string) (types.Account, error) {
	accountURL := cfg.API.BaseURL.JoinPath(cfg.API.AccountPath)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	body, err := services.GetRequest(
		accountURL,
		&headers,
	)
	if err != nil {
		return types.Account{}, fmt.Errorf("failed to fetch account information: %w", err)
	}

	var account types.Account
	err = json.Unmarshal(body, &account)
	if err != nil {
		return types.Account{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return account, nil
}

func SelectProfile(cfg config.Config, accessToken string, profileId string) (string, error) {
	profileURL := cfg.API.BaseURL.JoinPath(cfg.API.ProfilePath, profileId)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	headers.Add("Content-Type", "application/json")

	data := map[string]interface{}{
		"pin": nil,
	}

	response, err := services.PostRequest(
		profileURL,
		&headers,
		data,
	)
	if err != nil {
		return "", fmt.Errorf("failed to select a profile: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	token, _ := result["jwt"].(string)

	return token, nil
}
