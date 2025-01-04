package services

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/DaanVervacke/strips.be-archiver/pkg/types"
)

func GenerateUUID() (string, error) {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}

func sanitizeName(name string) string {
	re := regexp.MustCompile(`[^\p{L}0-9_]+`)
	reSpaces := regexp.MustCompile(`\s{2,}`)

	sanitizedName := re.ReplaceAllString(name, " ")
	sanitizedName = reSpaces.ReplaceAllString(sanitizedName, " ")
	sanitizedName = strings.TrimSpace(sanitizedName)
	sanitizedName = strings.ReplaceAll(sanitizedName, " ", "_")

	return sanitizedName
}

func IsValidUUID(u string) bool {
	uuidRegex := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`
	r := regexp.MustCompile(uuidRegex)
	return r.MatchString(u)
}

func CreateFileName(albumInfo types.Album) (string, error) {
	publicationDate, err := time.Parse("01/02/2006", albumInfo.PublicationDate)
	if err != nil {
		return "", fmt.Errorf("failed to parse publication date '%s': %w", albumInfo.PublicationDate, err)
	}

	seriesName := sanitizeName(albumInfo.Series.Name)
	capitalTitle := cases.Title(language.English).String(albumInfo.Title)
	title := sanitizeName(capitalTitle)

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

func MergeHeaders(h1, h2 http.Header) http.Header {
	merged := h1.Clone()
	for key, values := range h2 {
		for _, value := range values {
			merged.Add(key, value)
		}
	}
	return merged
}

func Cleanup(tempDir string) error {
	err := os.RemoveAll(tempDir)
	if err != nil {
		return fmt.Errorf("failed to remove temp directory")
	}

	return nil
}
