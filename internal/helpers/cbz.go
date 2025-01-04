package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func PrepareComicInfo(albumInfo types.Album) (types.ComicInfo, error) {
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
