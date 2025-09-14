package embed_env_test

import (
	"os"
	"testing"

	"github.com/digiconvent/embed_env"
)

func TestMain(t *testing.T) {
	thisBinary, _ := os.Executable()
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

	err := embed_env.ReadFromBinary(thisBinary, testInstance, "")
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

	err = embed_env.ReadFromBinary(thisBinary, testInstance2, "")
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

	err = embed_env.ReadFromBinary(thisBinary, testInstance2, "name=Annie+Hughes&age=38&male=false")
	if err != nil {
		t.Fatal("did not expect err, instead got", err.Error())
	}

	if testInstance2.Name != "Annie Hughes" {
		t.Fatal("expected name to be Annie Hughes, instead got", testInstance2.Name)
	}
	if testInstance2.Age != 38 {
		t.Fatal("expected age to be 38, instead got", testInstance2.Age)
	}
	if testInstance2.Male {
		t.Fatal("expected male to be false, instead got", testInstance2.Male)
	}
	t.Log("Passed [read from preset]")
}
