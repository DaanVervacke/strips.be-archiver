package types

import (
	"go/types"
	"net/url"
)

type Album struct {
	ProductID               types.Nil   `json:"productId"`
	EAN                     string      `json:"ean"`
	Language                string      `json:"language"`
	AgeRating               int         `json:"ageRating"`
	PublicationDate         string      `json:"publicationDate"`
	Authors                 []string    `json:"authors"`
	Illustrators            []string    `json:"illustrators"`
	Genres                  []string    `json:"genres"`
	Publisher               string      `json:"publisher"`
	Summary                 string      `json:"summary"`
	StatusForProfile        string      `json:"statusForProfile"`
	ID                      string      `json:"id"`
	Title                   string      `json:"title"`
	Series                  AlbumSeries `json:"series"`
	Thumbnail               string      `json:"thumbnail"`
	HighResolutionThumbnail string      `json:"highResolutionThumbnail"`
	Number                  int         `json:"sequence"`
	AmountOfPages           int
}

type AlbumSeries struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AlbumContent struct {
	URI string `json:"uri"`
}

type Playbook struct {
	LogicalBooks []Book `json:"logicalBooks"`
}

type Book struct {
	Assets Asset `json:"assets"`
}

type Asset struct {
	Images []Image `json:"images"`
}

type Image struct {
	Reference string `json:"reference"`
	Path      string `json:"path"`
	URL       *url.URL
}
