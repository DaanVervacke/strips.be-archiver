package handlers

import (
	"fmt"
	"log/slog"

	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func HandleLogin(email string, cfg config.Config) error {
	err := api.PostUserData(cfg, email)
	if err != nil {
		return err
	}

	slog.Info("user data has been posted, check your inbox and please enter your OTP code: ")

	var otp string
	_, err = fmt.Scanln(&otp)
	if err != nil {
		return err
	}

	supabaseAccessToken, err := api.VerifyUser(cfg, email, otp)
	if err != nil {
		return err
	}

	slog.Info("your account has been verified")

	stripsBeRefreshToken, deviceID, err := api.TradeJWT(cfg, supabaseAccessToken)
	if err != nil {
		return err
	}

	slog.Info("found supabase access token")

	stripsBeAccessToken, err := ProfileHandler(cfg, stripsBeRefreshToken)
	if err != nil {
		return err
	}

	slog.Info("login flow completed", "access_token", stripsBeAccessToken, "refresh_token", stripsBeRefreshToken, "device_id", deviceID)

	return nil
}
