package handlers

import (
	"math/rand"
	"sort"
	"time"

	"github.com/DaanVervacke/strips.be-archiver/internal/types"
	"github.com/DaanVervacke/strips.be-archiver/pkg/api"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
)

func HandleSeries(cfg config.Config, seriesID string, connections int, startAlbum int, endAlbum int, randomOrder bool, randomDelay bool, excludeMetadata bool) error {
	var albumNumbers []int

	if startAlbum > 0 && endAlbum > 0 && endAlbum >= startAlbum {
		albumNumbers = make([]int, 0, (endAlbum+1)-startAlbum)
		for i := startAlbum; i <= endAlbum; i++ {
			albumNumbers = append(albumNumbers, i)
		}
	}

	seriesInformation, err := api.GetSeriesInformation(cfg, seriesID)
	if err != nil {
		return err
	}

	if randomOrder {
		rand.Shuffle(len(seriesInformation.Albums), func(i, j int) {
			seriesInformation.Albums[i], seriesInformation.Albums[j] = seriesInformation.Albums[j], seriesInformation.Albums[i]
		})
	} else {
		sort.Slice(seriesInformation.Albums, func(i, j int) bool {
			return seriesInformation.Albums[i].Number < seriesInformation.Albums[j].Number
		})
	}

	albums := seriesInformation.Albums

	if len(albumNumbers) > 0 {
		albumNumbersMap := make(map[int]struct{})
		for _, num := range albumNumbers {
			albumNumbersMap[num] = struct{}{}
		}

		var filteredAlbums []types.Album
		for _, album := range seriesInformation.Albums {
			if _, ok := albumNumbersMap[album.Number]; ok {
				filteredAlbums = append(filteredAlbums, album)
			}
		}
		albums = filteredAlbums
	}

	for i, album := range albums {
		err := HandleAlbum(cfg, album.ID, connections, excludeMetadata)
		if err != nil {
			return err
		}

		if randomDelay && i < len(albums)-1 {
			pause := rand.Intn(31) + 30
			time.Sleep(time.Duration(pause) * time.Second)
		}
	}

	return nil
}
