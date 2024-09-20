package views

import (
	_ "embed"
	"encoding/json"
)

type LayoutConfig struct {
	CSSFileName, JSFileName string
}

//go:embed web/assets/dist/manifest.json
var manifestJSON []byte
var manifest = loadManifest()

// loadConfigData reads the embedded manifest.json file and returns a map of asset paths
func loadManifest() *LayoutConfig {
	manifest := make(map[string]string)
	err := json.Unmarshal(manifestJSON, &manifest)
	if err != nil {
		panic(err) // Don't run the app without the manifest
	}

	return &LayoutConfig{CSSFileName: manifest["main.css"], JSFileName: manifest["main.js"]}
}
