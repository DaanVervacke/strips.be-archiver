package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/config"
	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/internal/types"
)

func GetSeriesInformation(cfg config.Config, seriesID string) (types.Series, error) {
	seriesURL := cfg.API.BaseURL.JoinPath(cfg.API.SeriesPath, seriesID)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	body, err := services.GetRequest(
		seriesURL,
		&headers,
	)
	if err != nil {
		return types.Series{}, fmt.Errorf("failed to fetch series information: %w", err)
	}

	var serieInformation types.Series
	err = json.Unmarshal(body, &serieInformation)
	if err != nil {
		return types.Series{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return serieInformation, nil
}
