package api

import (
	"encoding/json"
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/helpers"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func GetSeriesInformation(cfg config.Config, seriesID string) (types.Series, error) {
	seriesURL := cfg.API.BaseURL.JoinPath(cfg.API.SeriesPath, seriesID)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	body, err := helpers.GetRequest(
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
