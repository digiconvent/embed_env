package embed_env_internal

import (
	"strings"
)

func ReadEmbeddedData(uri, delimiter string) (string, error) {
	lastLine, err := getLastLine(uri)
	if err != nil {
		return "", err
	}

	if strings.Count(lastLine, delimiter) == 0 {
		return "", nil
	}

	if strings.Contains(lastLine, delimiter) {
		segments := strings.Split(lastLine, delimiter)
		return segments[len(segments)-1], nil
	}

	return "", nil
}

