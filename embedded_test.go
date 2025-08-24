package embed_env_test

import (
	"testing"

	"github.com/DigiConvent/embed_env"
)

type someType struct {
	A string `name:"A"`
	B string `name:"xyz"`
	C int    `name:"age"`
	D bool   `name:"really"`
}

func TestGetRawEmbeddedData(t *testing.T) {
	a := someType{A: "text", B: "test2", C: -23, D: true}
	result, err := embed_env.ReadEmbeddedData(embed_env.Uri(), embed_env.Delimiter)
	if err != nil {
		t.Log(err)
	}
	if result != "" {
		t.Fatal("Expected no embedded data")
	}

	embed_env.WriteEmbeddedData(embed_env.Uri(), embed_env.Delimiter, a)

	result, err = embed_env.ReadEmbeddedData(embed_env.Uri(), embed_env.Delimiter)
	if err != nil {
		t.Log(err)
	}
	if result == "" {
		t.Fatal("Expected embedded data")
	}
}
