package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func temporaryFile() (*os.File, func()) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	return tmpFile, func() { os.Remove(tmpFile.Name()) }
}

func writePerson(person Person, sink io.WriteCloser) {
	err := json.NewEncoder(sink).Encode(person)
	if err != nil {
		panic(err)
	}
}

func readPerson(person *Person, source io.Reader) {
	err := json.NewDecoder(source).Decode(person)
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpFile, closer := temporaryFile()
	defer closer()

	toFile := Person{
		Name: "Kyle",
		Age:  27,
	}
	writePerson(toFile, tmpFile)
	tmpFile.Close()

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromFile Person
	readPerson(&fromFile, tmpFile2)

	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(toFile, fromFile) {
		fmt.Println("Failed!")
	} else {
		fmt.Println("Success!")
	}
}
