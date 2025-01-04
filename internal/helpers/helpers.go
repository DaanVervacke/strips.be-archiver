package helpers

import (
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"strings"
)

func GenerateUUID() (string, error) {
	UUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return UUID.String(), nil
}

func SanitizeName(name string) string {
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

func MergeHeaders(h1, h2 http.Header) http.Header {
	merged := h1.Clone()
	for key, values := range h2 {
		for _, value := range values {
			merged.Add(key, value)
		}
	}
	return merged
}
