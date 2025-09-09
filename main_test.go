package embed_env_test

import (
	"testing"

	"github.com/digiconvent/embed_env"
)

func TestMain(t *testing.T) {
	// make up some struct with data
	type SomeStruct struct {
		Name string `name:"name"`
		Age  int    `name:"age"`
		Male bool   `name:"male"`
	}

	testInstance := &SomeStruct{
		Name: "Dean McCoppin",
		Age:  40,
		Male: true,
	}

	err := embed_env.ReadFromBinary(testInstance)
	if err != nil {
		t.Fatal(err)
	}
	if testInstance.Name != "Dean McCoppin" || testInstance.Age != 40 || testInstance.Male != true {
		t.Fatal("did not expect values to be changed by a binary that has no embedded data")
	}
	t.Log("Passed [don't overwrite existing data]")

	err = embed_env.WriteToBinary(testInstance)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Passed [write data]")

	var testInstance2 = &SomeStruct{
		Name: "",
		Age:  -23,
		Male: false,
	}

	err = embed_env.ReadFromBinary(testInstance2)
	if err != nil {
		t.Fatal(err)
	}
	if testInstance.Name != testInstance2.Name && testInstance.Age == testInstance2.Age && testInstance.Male == testInstance2.Male {
		t.Fatal("expected", testInstance2, "to be", testInstance)
	}
	if testInstance == testInstance2 {
		t.Fatal("should not be the same instance")
	}
	t.Log("Passed [read existing data]")
}
