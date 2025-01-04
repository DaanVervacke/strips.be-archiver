package helpers

import (
	"bytes"
	"github.com/DaanVervacke/strips.be-archiver/pkg/config"
	"io"
	"net/url"
	"os"
)

func DownloadFile(cfg config.Config, filepath string, url *url.URL) error {
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
