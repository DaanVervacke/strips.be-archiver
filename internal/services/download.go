package services

import (
	"bytes"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/DaanVervacke/strips.be-archiver/internal/config"
	"github.com/DaanVervacke/strips.be-archiver/internal/types"
)

func DownloadFiles(cfg config.Config, images []types.Image, tempDir string, outputNameBase string, connections int) {
	var wg sync.WaitGroup
	slots := make(chan struct{}, connections)

	for _, image := range images {
		wg.Add(1)

		go func(img types.Image) {
			slots <- struct{}{}

			defer wg.Done()

			defer func() { <-slots }()

			err := downloadFile(cfg, filepath.Join(tempDir, outputNameBase+"_"+path.Base(img.Path)), img.URL)
			if err != nil {
				log.Fatal("Error downloading file:", err)
				return
			}
		}(image)
	}

	wg.Wait()
}

func downloadFile(cfg config.Config, filepath string, url *url.URL) error {
	headers := cfg.API.PlaybookHeaders.Clone()
	headers.Del("AppVersion")

	body, err := GetRequest(
		url,
		&headers,
	)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(body)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	var closeError error
	defer func() {
		if cerr := out.Close(); cerr != nil {
			closeError = cerr
		}
	}()

	_, err = io.Copy(out, reader)

	if closeError != nil {
		return closeError
	}

	return err
}
