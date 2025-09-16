package embed_env_internal

import (
	"os"
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

// start at the end of a file, load contents as long until a \n is found
func getLastLine(uri string) (string, error) {
	file, err := os.Open(uri)
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	var lastLine string
	var pos int64 = stat.Size()

	for {
		toRead := min(pos, int64(len(buf)))
		pos -= toRead

		_, err = file.ReadAt(buf[:toRead], pos)
		if err != nil {
			return "", err
		}

		for i := toRead - 1; i >= 0; i-- {
			if buf[i] == '\n' {
				if lastLine != "" {
					return lastLine, nil
				}
				continue
			}
			lastLine = string(buf[i]) + lastLine
		}

		if pos == 0 {
			return lastLine, nil
		}
	}
}
