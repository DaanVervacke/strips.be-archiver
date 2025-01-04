package services

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func CreateCBZ(tempPath, cbzFileName string, excludeMetadata bool) error {
	var closeError error

	zipfile, err := os.Create(cbzFileName + ".cbz")
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

	tempFiles, err := os.ReadDir(tempPath)
	if err != nil {
		return err
	}

	if !excludeMetadata {
		sort.SliceStable(tempFiles, func(i, j int) bool {
			return tempFiles[i].Name() == "ComicInfo.xml"
		})
	}

	for _, file := range tempFiles {
		filePath := filepath.Join(tempPath, file.Name())

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

func CreateComicInfoXML(albumInfo types.Album, tempPath string) error {
	comicInfo, err := prepareComicInfo(albumInfo)
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

func prepareComicInfo(albumInfo types.Album) (types.ComicInfo, error) {
	comicInfo := types.ComicInfo{}

	publicationDate, err := time.Parse("01/02/2006", albumInfo.PublicationDate)
	if err != nil {
		return types.ComicInfo{}, fmt.Errorf("failed to parse publication date: %v", err)
	}

	if albumInfo.Title != "" {
		comicInfo.Title = albumInfo.Title
	}

	if albumInfo.Series != (types.AlbumSeries{}) && albumInfo.Series.Name != "" {
		comicInfo.Series = albumInfo.Series.Name
	}

	if albumInfo.Number > 0 {
		comicInfo.Number = albumInfo.Number
	}

	if albumInfo.Summary != "" {
		comicInfo.Summary = albumInfo.Summary
	}

	comicInfo.Year = publicationDate.Year()
	comicInfo.Month = int(publicationDate.Month())
	comicInfo.Day = publicationDate.Day()

	if authors := strings.Join(albumInfo.Authors, ", "); authors != "" {
		comicInfo.Writer = authors
	}

	if illustrators := strings.Join(albumInfo.Illustrators, ", "); illustrators != "" {
		comicInfo.Penciller = illustrators
	}

	if albumInfo.Publisher != "" {
		comicInfo.Publisher = albumInfo.Publisher
	}

	if genres := strings.Join(albumInfo.Genres, ", "); genres != "" {
		comicInfo.Genre = genres
	}

	if albumInfo.AmountOfPages > 0 {
		comicInfo.PageCount = albumInfo.AmountOfPages
	}

	if language := strings.ToLower(albumInfo.Language); language != "" {
		comicInfo.LanguageISO = language
	}

	comicInfo.Format = "Digital"

	if albumInfo.AgeRating > -1 {
		comicInfo.AgeRating = albumInfo.AgeRating
	}

	if albumInfo.EAN != "" {
		comicInfo.GTIN = albumInfo.EAN
	}

	return comicInfo, nil
}
