package services

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/DaanVervacke/strips.be-archiver/internal/helpers"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func CreateFileName(albumInfo types.Album) (string, error) {
	publicationDate, err := time.Parse("01/02/2006", albumInfo.PublicationDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse publication date '%s': %w", albumInfo.PublicationDate, err)
	}

	seriesName := helpers.SanitizeName(albumInfo.Series.Name)
	capitalTitle := cases.Title(language.English).String(albumInfo.Title)
	title := helpers.SanitizeName(capitalTitle)

	sequence := albumInfo.Number
	year := strconv.Itoa(publicationDate.Year())

	fileName := fmt.Sprintf("%s_%s_%d_%s", seriesName, title, sequence, year)
	fileName = strings.TrimSpace(fileName)

	return fileName, nil
}

func CreateTempDir(currentDir string) (string, error) {
	tempDir := filepath.Join(currentDir, "temp")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("failed to create temp directory %q: %v", tempDir, err)
		}
		return "", fmt.Errorf("unexpected error creating temp directory %q: %v", tempDir, err)
	}
	return tempDir, nil
}

func Cleanup(tempDir string) error {
	err := os.RemoveAll(tempDir)
	if err != nil {
		return fmt.Errorf("failed to remove temp directory")
	}

	return nil
}

func CheckDir(outputDir string) error {
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return fmt.Errorf("invalid output directory %q", outputDir)
	}

	return nil
}

func DownloadImages(cfg config.Config, images []types.Image, tempDir string, outputNameBase string, connections int) error {
	var wg sync.WaitGroup
	var downloadErr error
	slots := make(chan struct{}, connections)

	for _, image := range images {
		wg.Add(1)

		go func(img types.Image) {
			slots <- struct{}{}

			defer wg.Done()

			defer func() { <-slots }()

			err := helpers.DownloadFile(cfg, filepath.Join(tempDir, outputNameBase+"_"+path.Base(img.Path)), img.URL)
			if err != nil {
				downloadErr = err
			}
		}(image)
	}

	wg.Wait()

	return downloadErr
}

func CreateComicInfoXML(albumInfo types.Album, tempPath string) error {
	comicInfo, err := helpers.PrepareComicInfo(albumInfo)
	if err != nil {
		return fmt.Errorf("failed to prepare comic info: %v", err)
	}

	comicInfoWrapper := types.ComicInfoWrapper{
		XSI:       "http://www.w3.org/2001/XMLSchema-instance",
		XSD:       "http://www.w3.org/2001/XMLSchema",
		ComicInfo: comicInfo,
	}

	xmlData, err := xml.MarshalIndent(comicInfoWrapper, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal comic info: %v", err)
	}

	xmlFile, err := os.Create(filepath.Join(tempPath, "ComicInfo.xml"))
	if err != nil {
		return fmt.Errorf("failed to create ComicInfo.xml")
	}

	var closeError error
	defer func() {
		if cerr := xmlFile.Close(); closeError != nil {
			closeError = cerr
		}
	}()

	_, err = xmlFile.WriteString(`<?xml version="1.0"?>` + "\n")
	if err != nil {
		return fmt.Errorf("failed to write XML version declaration to ComicInfo.xml")
	}

	_, err = xmlFile.Write(xmlData)
	if err != nil {
		return fmt.Errorf("failed to write data to ComicInfo.xml")
	}

	if closeError != nil {
		return closeError
	}

	return nil
}

func CreateCBZ(outputDir string, tempDir string, fileName string, excludeMetadata bool) error {
	var closeError error

	outputPath := filepath.Join(outputDir, fileName+".cbz")

	zipfile, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := zipfile.Close(); cerr != nil {
			closeError = cerr
		}
	}()

	zipWriter := zip.NewWriter(zipfile)

	defer func() {
		if cerr := zipWriter.Close(); cerr != nil {
			closeError = cerr
		}
	}()

	tempFiles, err := os.ReadDir(tempDir)
	if err != nil {
		return err
	}

	if !excludeMetadata {
		sort.SliceStable(tempFiles, func(i, j int) bool {
			return tempFiles[i].Name() == "ComicInfo.xml"
		})
	}

	for _, file := range tempFiles {
		filePath := filepath.Join(tempDir, file.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		info, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Store

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = writer.Write(data)
		if err != nil {
			return err
		}
	}

	if closeError != nil {
		return closeError
	}

	return nil
}
