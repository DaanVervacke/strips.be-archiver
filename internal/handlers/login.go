package handlers

import (
	"fmt"

	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func HandleLogin(email string, cfg config.Config) error {
	err := api.PostUserData(cfg, email)
	if err != nil {
		return err
	}

	fmt.Printf("%s User data has been posted. Check your inbox and please enter your OTP code: ", services.SuccessStyle.Render("SUCCESS"))

	var otp string
	_, err = fmt.Scanln(&otp)
	if err != nil {
		return err
	}

	supabaseAccessToken, err := api.VerifyUser(cfg, email, otp)
	if err != nil {
		return err
	}

	fmt.Println(services.SuccessStyle.Render("SUCCESS"), "Your account has been verified!")

	stripsBeRefreshToken, deviceID, err := api.TradeJWT(cfg, supabaseAccessToken)
	if err != nil {
		return err
	}

	fmt.Println(services.SuccessStyle.Render("SUCCESS"), "Found Supabase access token!")

	stripsBeAccessToken, err := ProfileHandler(cfg, stripsBeRefreshToken)
	if err != nil {
		return err
	}

	fmt.Printf("%s Login flow has been completed!\n\nYour access token is: %s\n\nYour refresh token is: %s\n\nYour device id is: %s\n", services.SuccessStyle.Render("SUCCESS"), stripsBeAccessToken, stripsBeRefreshToken, deviceID)

	return nil
}
