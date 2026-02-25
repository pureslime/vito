package engine

import (
	"bytes"
	"context"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func ExecuteWasmPill(wasmPath string, args []string) error {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	wasmCode, err := os.ReadFile(wasmPath)
	if err != nil {
		return err
	}

	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithStdin(os.Stdin).
		WithArgs(append([]string{wasmPath}, args...)...)

	_, err = r.InstantiateWithConfig(ctx, wasmCode, config)
	return err
}

func GetWasmOutput(path string, arg string) ([]byte, error) {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	wasmCode, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	config := wazero.NewModuleConfig().
		WithStdout(&buf).
		WithArgs("vito-pill", arg)

	_, err = r.InstantiateWithConfig(ctx, wasmCode, config)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
