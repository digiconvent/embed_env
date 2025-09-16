package embed_env

import (
	embed_env_internal "github.com/digiconvent/embed_env/internal"
)

// pass a pointer to your struct struct and name the individual fields using tags like `name:"some_env_name"`
// fields should only use primitive values
func ReadFromBinary(binaryUri string, someStruct any, preset string) error {
	var err error
	var query string = preset
	if query == "" {
		query, err = ReadEmbeddedData(binaryUri)
		if err != nil {
			return err
		}
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

// binary uri to scan for embedded data. This function can be used to scan the embedded data from another digiconvent binary
func ReadEmbeddedData(binPath string) (string, error) {
	query, err := embed_env_internal.ReadEmbeddedData(binPath, embed_env_internal.Delimiter)
	if err != nil {
		return "", err
	}
	return query, nil
}
