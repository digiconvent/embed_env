package embed_env_internal

import (
	"bytes"
	"os"
)

// pray to the Lord that this string does not occur
const Delimiter string = "Z0E8LSlyfX12P3RDX3JHLTU4Tno1MjtQ"

func GetDelimiterPositions(delimiter string) ([]int, error) {
	data, err := os.ReadFile(Uri())
	if err != nil {
		return nil, err
	}

	var locations []int
	start := 0
	for {
		idx := bytes.Index(data[start:], []byte(delimiter))
		if idx == -1 {
			break
		}
		locations = append(locations, start+idx)
		start += idx + 1
	}
	return locations, nil
}
