package handlers

import (
	"log/slog"

	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func HandleAlbumsSearch(cfg config.Config, searchKeyword string) error {
	albumResults, err := api.SearchAlbums(cfg, searchKeyword)
	if err != nil {
		return err
	}

	if len(albumResults) > 10 {
		albumResults = albumResults[:10]
	}

	for _, album := range albumResults {
		slog.Info("found album", "title", album.Title, "series", album.Series.Name, "id", album.ID)
	}

	return nil
}

func HandleSeriesSearch(cfg config.Config, searchKeyword string) error {
	results, err := api.SearchSeries(cfg, searchKeyword)
	if err != nil {
		return err
	}

	if len(results) > 10 {
		results = results[:10]
	}

	for _, result := range results {
		slog.Info("found series", "name", result.Name, "id", result.ID)
	}

	return nil
}
