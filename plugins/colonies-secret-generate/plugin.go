package main

import (
	pdk "github.com/extism/go-pdk"

	"github.com/colonyos/colonies/pkg/security/crypto"
)

//export prvkey
func prvkey() int32 {
	prvKey, err := crypto.CreateCrypto().GeneratePrivateKey()
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString(prvKey)
	return 0
}

//export id
func id() int32 {
	inputData := pdk.Input()
	prvKey := string(inputData)

	id, err := crypto.CreateCrypto().GenerateID(prvKey)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString(id)
	return 0
}

func main() {}
