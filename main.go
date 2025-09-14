package embed_env

import (
	"fmt"
	"os"

	embed_env_internal "github.com/digiconvent/embed_env/internal"
)

// pass a pointer to your struct struct and name the individual fields using tags like `name:"some_env_name"`
// fields should only use primitive values
func ReadFromBinary(someStruct any, preset string) error {
	thisBinary, err := os.Executable()
	if err != nil {
		return err
	}
	var query string = preset
	if query == "" {
		query, err = embed_env_internal.ReadEmbeddedData(thisBinary, embed_env_internal.Delimiter)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("using preset ", preset)
	}

	if query == "" {
		return nil
	}

	return embed_env_internal.FromQuery(someStruct, &query)
}

// this parameter should be a pointer to the struct which contains the data which you want to embed
func WriteToBinary(someStruct any) error {
	return embed_env_internal.WriteEmbeddedData(embed_env_internal.Uri(), embed_env_internal.Delimiter, someStruct)
}
