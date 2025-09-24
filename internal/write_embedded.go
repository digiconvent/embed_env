package embed_env_internal

import (
	"errors"
	"os"
	"strings"
)

func WriteEmbeddedData(uri, delimiter string, someStruct any) error {
	lastLine, err := getLastLine(uri)
	if err != nil {
		return err
	}

	hasEmbeddedData := strings.Contains(lastLine, delimiter)

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

	var newContent []byte
	if hasEmbeddedData {
		delimiterIndex := strings.Index(lastLine, delimiter)
		if delimiterIndex != -1 {
			bytesToRemove := len(lastLine) - delimiterIndex
			newContent = content[:len(content)-bytesToRemove]
		} else {
			newContent = content
		}
		newContent = append(newContent, []byte(delimiter+data)...)
	} else {
		newContent = append(content, []byte(delimiter+data)...)
	}

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
