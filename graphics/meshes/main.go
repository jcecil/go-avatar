package meshes

//go:generate go-bindata -debug -o meshes.go -pkg meshes .

import (
	"bytes"
	"fmt"
	"github.com/GlenKelley/go-collada"
)

func LoadColladaCube(assetName string) (uint32, error) {
	data, err := Asset(assetName)
	if err != nil {
		return 0, err
	}

	colladaDoc, err := collada.LoadDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	fmt.Println("Collada doc: ", colladaDoc)
	return 0, err
}
