package main

import (
	"encoding/json"

	pdk "github.com/extism/go-pdk"

	"github.com/colonyos/colonies/pkg/security/crypto"
	"gopkg.in/yaml.v3"
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

// replaceColoniesPrivateKey recursively traverses the YAML structure
// and replaces the value of any env variable with name "COLONIES_PRVKEY"
func replaceColoniesPrivateKey(node interface{}, privateKey string) {
	switch v := node.(type) {
	case map[string]interface{}:
		// Check if this is an env variable object with name COLONIES_PRVKEY
		if name, ok := v["name"].(string); ok && name == "COLONIES_PRVKEY" {
			v["value"] = privateKey
		}
		// Recursively process all values in the map
		for _, value := range v {
			replaceColoniesPrivateKey(value, privateKey)
		}
	case []interface{}:
		// Recursively process all items in the array
		for _, item := range v {
			replaceColoniesPrivateKey(item, privateKey)
		}
	}
}

//export helm_plugin_main
func helm_plugin_main() int32 {
	// Define the plugin input/output structure
	type PluginData struct {
		Manifests map[string]string `json:"manifests"`
		ExtraArgs []string          `json:"extraArgs"`
	}

	// Read JSON input from stdin
	inputData := pdk.Input()

	// Generate private key
	prvKey, err := crypto.CreateCrypto().GeneratePrivateKey()
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	// Parse JSON
	var pluginData PluginData
	err = json.Unmarshal(inputData, &pluginData)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	// Process each manifest
	for filename, manifestYAML := range pluginData.Manifests {
		// Parse the manifest YAML
		var data interface{}
		err = yaml.Unmarshal([]byte(manifestYAML), &data)
		if err != nil {
			pdk.SetError(err)
			return 1
		}

		// Replace all COLONIES_PRVKEY values with the generated private key
		replaceColoniesPrivateKey(data, prvKey)

		// Marshal back to YAML
		modifiedYAML, err := yaml.Marshal(data)
		if err != nil {
			pdk.SetError(err)
			return 1
		}

		// Update the manifest in the map
		pluginData.Manifests[filename] = string(modifiedYAML)
	}

	// Marshal back to JSON
	output, err := json.Marshal(pluginData)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	// Output the modified JSON
	pdk.Output(output)
	return 0
}

func main() {}
