package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DaanVervacke/strips.be-archiver/internal/config"
	"github.com/DaanVervacke/strips.be-archiver/internal/handlers"
	"github.com/DaanVervacke/strips.be-archiver/internal/services"
)

func main() {

	var err error

	albumIDFlag := flag.String("download-album", "", "The ID of the album to be processed. This should be a valid UUID.")
	seriesIDFlag := flag.String("download-series", "", "The ID of the series to be processed. This should be a valid UUID.")

	searchAlbumsFlag := flag.String("search-album", "", "An album name to search for. Returns matching IDs.")
	searchSeriesFlag := flag.String("search-series", "", "A series name to search for. Returns matching IDs.")

	configFlag := flag.String("config", "", "Path to your strips.be config file")

	loginFlag := flag.String("login", "", "The e-mail address associated with your strips.be account.")
	refreshFlag := flag.Bool("refresh", false, "Refresh the access token that is currently associated with your strips.be account.")

	connectionsFlag := flag.Int("connections", 1,
		"The amount of simultaneous connections this tool will make to the strips.be api.\nThis only applies to the download flags.\nToo many connections WILL get you blacklisted.")

	rangeStartFlag := flag.Int("album-start", 0, "The number of the first album to be downloaded. Only applies when downloading a series.")
	rangeEndFlag := flag.Int("album-end", 0, "The number of the last album to be downloaded. Only applies when downloading a series.")

	randomizeOrderFlag := flag.Bool("randomize-order", false, "Use this flag to randomize the download order of a series.")
	randomizeDelayFlag := flag.Bool("randomize-delay", false, "Use this flag to randomize the delay between each download.")

	excludeMetadataFlag := flag.Bool("exclude-metadata", false, "Use this flag to exclude the ComicInfo.xml file from the final archive.")

	flag.Parse()

	if *configFlag == "" {
		err := fmt.Errorf("config file not set, please create one and use the --config flag to parse it")
		handleError(err)
	}

	cfg, err := config.LoadConfig(*configFlag)
	handleError(err)

	if *albumIDFlag == "" && *seriesIDFlag == "" && *searchAlbumsFlag == "" && *searchSeriesFlag == "" && *loginFlag == "" && !*refreshFlag {
		err := fmt.Errorf("either the 'download-album', 'download-series', 'search-album', 'search-series', 'login' or 'refresh' flag is required. Please provide a valid flag")
		handleError(err)
	}

	if *albumIDFlag != "" && !services.IsValidUUID(*albumIDFlag) {
		err := fmt.Errorf("invalid album ID format. Expected format is a valid UUID")
		handleError(err)
	}

	if *seriesIDFlag != "" && !services.IsValidUUID(*seriesIDFlag) {
		err := fmt.Errorf("invalid series ID format. Expected format is a valid UUID")
		handleError(err)
	}

	if *loginFlag != "" {
		err = handlers.HandleLogin(*loginFlag, cfg)
	} else if *refreshFlag {
		err = handlers.HandleRefresh(cfg)
	} else if *albumIDFlag != "" {
		err = handlers.HandleAlbum(cfg, *albumIDFlag, *connectionsFlag, *excludeMetadataFlag)
	} else if *seriesIDFlag != "" {
		err = handlers.HandleSeries(cfg, *seriesIDFlag, *connectionsFlag, *rangeStartFlag, *rangeEndFlag, *randomizeOrderFlag, *randomizeDelayFlag, *excludeMetadataFlag)
	} else if *searchAlbumsFlag != "" {
		err = handlers.HandleAlbumsSearch(cfg, *searchAlbumsFlag)
	} else if *searchSeriesFlag != "" {
		err = handlers.HandleSeriesSearch(cfg, *searchSeriesFlag)
	}

	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%s %v\n", services.ErrorStyle.Render("ERROR"), err.Error())
		os.Exit(1)
	}
}
