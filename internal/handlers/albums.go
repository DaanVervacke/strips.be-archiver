package handlers

import (
	"fmt"
	"os"

	"github.com/DaanVervacke/strips.be-archiver/internal/config"
	"github.com/DaanVervacke/strips.be-archiver/internal/services"
	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
)

func HandleAlbum(cfg config.Config, albumID string, connections int, excludeMetadata bool) error {
	albumInformation, err := api.GetAlbumInformation(cfg, albumID)
	if err != nil {
		return err
	}

	if albumInformation.StatusForProfile != "AVAILABLE" {
		fmt.Println(services.WarningStyle.Render("WARNING"), fmt.Sprintf("album %s is not available for your profile, skipping...", albumID))
		return nil
	}

	outputName, err := services.CreateFileName(albumInformation)
	if err != nil {
		return err
	}

	fmt.Println(services.TitleStyle.Render(" ALBUM INFO  "), fmt.Sprintf("Title: %s | Series: %s | Release Date: %s", albumInformation.Title, albumInformation.Series.Name, albumInformation.PublicationDate))
	fmt.Println(services.TitleStyle.Render(" EXPORT INFO "), fmt.Sprintf("Output Name: %s", outputName))

	playbookURL, err := api.GetPlaybookURL(cfg, albumID)
	if err != nil {
		return err
	}

	fmt.Println(services.SuccessStyle.Render("SUCCESS"), "The playbook URL has been found.")

	playbookContent, err := api.GetPlaybookContent(cfg, playbookURL)
	if err != nil {
		return err
	}

	fmt.Println(services.SuccessStyle.Render("SUCCESS"), "The playbook content has been parsed.")

	albumInformation.AmountOfPages = len(playbookContent)

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	tempDir, err := services.CreateTempDir(currentDir)
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}

	if !excludeMetadata {
		err = services.CreateComicInfoXML(albumInformation, tempDir)
		if err != nil {
			return err
		}

		fmt.Println(services.SuccessStyle.Render("SUCCESS"), "The metadata file related to this album has been created.")
	}

	services.DownloadFiles(cfg, playbookContent, tempDir, outputName, connections)

	fmt.Println(services.SuccessStyle.Render("SUCCESS"), "The images related to this album have been downloaded.")

	err = services.CreateCBZ(tempDir, outputName, excludeMetadata)
	if err != nil {
		return err
	}

	fmt.Println(services.SuccessStyle.Render("SUCCESS") + " Your archive has been created and saved as: " + outputName + ".cbz")

	err = services.Cleanup(tempDir)
	if err != nil {
		return err
	}

	return nil
}
