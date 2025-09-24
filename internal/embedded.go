package embed_env_internal

import (
	"os"
)

// pray to the Lord that this string does not occur
const Delimiter string = "Z0E8LSlyfX12P3RDX3JHLTU4Tno1MjtQ"

// getLastLine reads the last line from a file efficiently
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

