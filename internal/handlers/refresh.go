package handlers

import (
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/helpers"
	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func HandleRefresh(cfg config.Config) error {
	if cfg.Auth.Account.RefreshToken == "" {
		return fmt.Errorf("error getting the strips.be refresh token")
	}

	if cfg.Auth.Account.DeviceID == "" {
		return fmt.Errorf("error getting the strips.be device id")
	}

	newRefreshToken, err := api.RefreshJWT(cfg)
	if err != nil {
		return err
	}

	newAccessToken, err := ProfileHandler(cfg, newRefreshToken)
	if err != nil {
		return err
	}

	fmt.Printf("%s Refresh flow has been completed!\n\nYour new access token is: %s\n\nYour new refresh token is: %s\n", helpers.SuccessStyle.Render("SUCCESS"), newAccessToken, newRefreshToken)

	return nil
}
