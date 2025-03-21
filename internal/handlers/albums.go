package handlers

import (
	"fmt"
	"log/slog"

	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/services"
)

func HandleAlbum(cfg config.Config, albumID string, connections int, excludeMetadata bool, outputDir string) error {
	albumInformation, err := api.GetAlbumInformation(cfg, albumID)
	if err != nil {
		return err
	}

	if albumInformation.StatusForProfile != "AVAILABLE" {
		slog.Warn("album is not available for your profile, skipping...", "id", albumID)
		return nil
	}

	outputName, err := services.CreateFileName(albumInformation)
	if err != nil {
		return err
	}

	slog.Info("album info", "title", albumInformation.Title, "series", albumInformation.Series.Name, "release date", albumInformation.PublicationDate)
	slog.Info("export info", "output name", outputName)

	playbookURL, err := api.GetPlaybookURL(cfg, albumID)
	if err != nil {
		return err
	}

	slog.Info("the playbook URL has been found", "url", playbookURL)

	playbookContent, err := api.GetPlaybookContent(cfg, playbookURL)
	if err != nil {
		return err
	}

	slog.Info("the playbook content has been parsed")

	albumInformation.AmountOfPages = len(playbookContent)

	tempDir, err := services.CreateTempDir(outputDir)
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}

	if !excludeMetadata {
		err = services.CreateComicInfoXML(albumInformation, tempDir)
		if err != nil {
			return err
		}

		slog.Info("the metadata file related to this album has been created")
	}

	if err := services.DownloadImages(cfg, playbookContent, tempDir, outputName, connections); err != nil {
		return fmt.Errorf("error while downloading images: %v", err)
	}

	slog.Info("the images related to this album have been downloaded.")

	err = services.CreateCBZ(outputDir, tempDir, outputName, excludeMetadata)
	if err != nil {
		return err
	}

	slog.Info("your archive has been created", "filename", outputName+".cbz")

	err = services.Cleanup(tempDir)
	if err != nil {
		return err
	}

	return nil
}
