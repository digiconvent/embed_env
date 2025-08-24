package embedded_test

import (
	"testing"

	"github.com/DigiConvent/embed_env/embedded"
)

type someType struct {
	A string `name:"A"`
	B string `name:"xyz"`
	C int    `name:"age"`
	D bool   `name:"really"`
}

func TestGetRawEmbeddedData(t *testing.T) {
	a := someType{A: "text", B: "test2", C: -23, D: true}
	result, err := embedded.ReadEmbeddedData(embedded.Uri(), embedded.Delimiter)
	if err != nil {
		t.Log(err)
	}
	if result != "" {
		t.Fatal("Expected no embedded data")
	}

	embedded.WriteEmbeddedData(embedded.Uri(), embedded.Delimiter, a)

	result, err = embedded.ReadEmbeddedData(embedded.Uri(), embedded.Delimiter)
	if err != nil {
		t.Log(err)
	}
	if result == "" {
		t.Fatal("Expected embedded data")
	}
	t.Log(result)
}
