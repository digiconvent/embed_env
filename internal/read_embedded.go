package embed_env_internal

import (
	"os"
)

func ReadEmbeddedData(uri, delimiter string) (string, error) {
	contents, err := os.ReadFile(uri)
	if err != nil {
		return "", err
	}

	pos, err := GetDelimiterPositions(delimiter)
	if err != nil {
		return "", err
	}

	if len(pos) != 2 { // there are 2 positions where the delimiter occurs if there is embedded data
		return "", nil
	}

	return string(contents[pos[1]+len(delimiter):]), nil
}
