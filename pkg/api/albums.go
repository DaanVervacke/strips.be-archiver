package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/DaanVervacke/strips.be-archiver/internal/helpers"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func GetAlbumInformation(cfg config.Config, albumID string) (types.Album, error) {
	albumURL := cfg.API.BaseURL.JoinPath(cfg.API.AlbumPath, albumID)

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	body, err := helpers.GetRequest(
		albumURL,
		&headers,
	)
	if err != nil {
		return types.Album{}, fmt.Errorf("failed to fetch album information: %w", err)
	}

	var albumInformation types.Album
	err = json.Unmarshal(body, &albumInformation)
	if err != nil {
		return types.Album{}, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return albumInformation, nil
}

func GetPlaybookURL(cfg config.Config, albumID string) (*url.URL, error) {
	albumContentURL := cfg.API.BaseURL.JoinPath(cfg.API.AlbumPath, albumID, "content")

	headers := cfg.API.BasicHeaders.Clone()
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", cfg.Auth.Account.AccessToken))

	body, err := helpers.GetRequest(albumContentURL, &headers)
	if err != nil {
		return nil, fmt.Errorf("failed to get playbook URL: %w", err)
	}

	var albumContent types.AlbumContent
	err = json.Unmarshal(body, &albumContent)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	playbookURL, err := url.Parse(albumContent.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri: %w", err)
	}

	playbookURL = playbookURL.JoinPath("playbook-classic.json")

	return playbookURL, nil
}

func GetPlaybookContent(cfg config.Config, playbookURL *url.URL) ([]types.Image, error) {
	body, err := helpers.GetRequest(playbookURL, &cfg.API.PlaybookHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed to get playbook content: %w", err)
	}

	var playbook types.Playbook
	err = json.Unmarshal(body, &playbook)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	images := playbook.LogicalBooks[0].Assets.Images

	for i := range images {
		newURL := *playbookURL
		basePath := path.Dir(playbookURL.Path)
		newURL.Path = path.Join(basePath, images[i].Path)
		images[i].URL = &newURL
	}

	return images, nil
}
