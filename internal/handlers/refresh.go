package handlers

import (
	"fmt"
	"log/slog"

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

	slog.Info("refresh flow has been completed", "access_token", newAccessToken, "refresh_token", newRefreshToken)

	return nil
}
