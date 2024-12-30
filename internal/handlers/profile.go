package handlers

import (
	"github.com/DaanVervacke/strips.be-archiver/internal/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
)

func ProfileHandler(cfg config.Config, stripsBeAccessToken string) (string, error) {
	accountInformation, err := api.GetAccount(cfg, stripsBeAccessToken)
	if err != nil {
		return "", err
	}

	profileID := accountInformation.Profiles[0].ID

	stripsBeProfileAccessToken, err := api.SelectProfile(cfg, stripsBeAccessToken, profileID)
	if err != nil {
		return "", err
	}

	return stripsBeProfileAccessToken, nil
}
