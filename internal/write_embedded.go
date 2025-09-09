package embed_env_internal

import (
	"errors"
	"os"
	"strings"
)

func WriteEmbeddedData(uri, delimiter string, someStruct any) error {
	locs, err := GetDelimiterPositions(delimiter)
	if err != nil {
		return err
	}

	marker := 0
	if len(locs) == 2 {
		marker = locs[1] + len(delimiter)
	}

	data, err := ToQuery(someStruct)
	if err != nil {
		return err
	}
	if strings.Contains(data, delimiter) {
		return errors.New("data contains " + delimiter + " and that is illegal")
	}

	content, err := os.ReadFile(uri)
	if err != nil {
		return err
	}
	if marker == 0 {
		marker = len(content)
		data = delimiter + data
	}

	base := content[:marker]

	newContent := append(base, []byte(data)...)

	// delete myself
	err = os.Remove(uri)
	if err != nil {
		return err
	}

	err = os.WriteFile(uri, newContent, 0755)
	if err != nil {
		return err
	}

	return nil
}
