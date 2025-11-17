package plugin_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	extism "github.com/extism/go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/wazero"
)

func init() {
	extism.SetLogLevel(extism.LogLevelDebug)
}

func loadFilePlugin(ctx context.Context) (*extism.Plugin, error) {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "../plugin.wasm",
				Name: "colonies-security-generate",
			},
			//extism.WasmData{
			//	Data: pluginBytes,
			//	Name: "gotemplate-renderer",
			//},
		},
		Memory: &extism.ManifestMemory{
			MaxPages: 65535,
			// MaxHttpResponseBytes: 1024 * 1024 * 10,
			// MaxVarBytes:          1024 * 1024 * 10,
		},
		Config:       map[string]string{},
		AllowedPaths: map[string]string{},
		Timeout:      0,
	}

	config := extism.PluginConfig{
		ModuleConfig:  wazero.NewModuleConfig().WithSysWalltime(),
		RuntimeConfig: wazero.NewRuntimeConfig().WithCloseOnContextDone(false),
		EnableWasi:    true,
		// EnableHttpResponseHeaders: true,
		// ObserveAdapter: ,
		// ObserveOptions: &observe.Options{},
	}

	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plugin: %w", err)
	}

	plugin.SetLogger(func(logLevel extism.LogLevel, s string) {
		fmt.Printf("%s %s: %s\n", time.Now().Format(time.RFC3339), logLevel.String(), s)
	})

	return plugin, nil
}

func TestIdDeterminism(t *testing.T) {
	ctx := t.Context()

	plugin, err := loadFilePlugin(ctx)
	require.NoError(t, err)
	defer plugin.Close(ctx)

	// First, generate a private key using the prvkey function
	exitCode, prvKeyData, err := plugin.Call("prvkey", nil)
	require.NoError(t, err)
	require.Equal(t, uint32(0), exitCode)
	prvKey := string(prvKeyData)
	require.NotEmpty(t, prvKey, "generated private key should not be empty")

	// Verify the private key is 64 hexadecimal characters
	require.Len(t, prvKey, 64, "private key should be exactly 64 characters")
	require.Regexp(t, "^[0-9a-fA-F]{64}$", prvKey,
		"private key should contain only hexadecimal characters")

	exitCode, idData, err := plugin.Call("id", []byte(prvKey))
	require.Equal(t, uint32(0), exitCode, "ID function gave an error code %d", exitCode)
	require.NoError(t, err)
	exitCode, idData2, err := plugin.Call("id", []byte(prvKey))
	require.Equal(t, uint32(0), exitCode, "ID function gave an error code %d", exitCode)
	require.NoError(t, err)

	// Verify all IDs are identical
	assert.Equal(t, idData, idData2,
		"ID function is non deterministic.\nTested with private key:%s\nGot IDs:\n%s\n%s", string(prvKeyData), string(idData), string(idData2))
}
