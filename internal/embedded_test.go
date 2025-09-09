package embed_env_internal_test

import (
	"testing"

	embed_env_internal "github.com/digiconvent/embed_env/internal"
)

const someTypeQuery = "age=-23&ax=text&really=true&xyz=test2"

type someType struct {
	A string `name:"ax"`  // this should be the second entry in the query since age < ax
	B string `name:"xyz"` // this should be the last entry in the query, obviously
	C int    `name:"age"`
	D bool   `name:"really"`
}

func TestGetRawEmbeddedData(t *testing.T) {
	a := someType{A: "text", B: "test2", C: -23, D: true}
	result, err := embed_env_internal.ReadEmbeddedData(embed_env_internal.Uri(), embed_env_internal.Delimiter)
	if err != nil {
		t.Log(err)
	}
	if result != "" {
		t.Fatal("Expected no embedded data")
	}

	embed_env_internal.WriteEmbeddedData(embed_env_internal.Uri(), embed_env_internal.Delimiter, a)

	result, err = embed_env_internal.ReadEmbeddedData(embed_env_internal.Uri(), embed_env_internal.Delimiter)
	if err != nil {
		t.Log(err)
	}

	if result != someTypeQuery {
		t.Fatal("Expected embedded data to be", someTypeQuery, "got", result, "instead")
	}
}
