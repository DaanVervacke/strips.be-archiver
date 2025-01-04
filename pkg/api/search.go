package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func SearchAlbums(cfg config.Config, contentKeyword string) ([]types.Album, error) {
	searchUrl := cfg.API.BaseURL.JoinPath(cfg.API.AlbumPath)

	params := url.Values{}
	params.Add("searchText", contentKeyword)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	searchUrl.RawQuery = params.Encode()

	body, err := services.GetRequest(searchUrl, &headers)
	if err != nil {
		return nil, fmt.Errorf("album search failed: %w", err)
	}

	var result struct {
		Content []types.Album `json:"albums"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return result.Content, nil
}

func SearchSeries(cfg config.Config, contentKeyword string) ([]types.Series, error) {
	searchUrl := cfg.API.BaseURL.JoinPath(cfg.API.SeriesPath)

	params := url.Values{}
	params.Add("searchText", contentKeyword)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	searchUrl.RawQuery = params.Encode()

	body, err := services.GetRequest(searchUrl, &headers)
	if err != nil {
		return nil, fmt.Errorf("series search failed: %w", err)
	}

	var result struct {
		Content []types.Series `json:"series"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return result.Content, nil
}
