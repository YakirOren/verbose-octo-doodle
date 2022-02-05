wasm:
	GOOS=js GOARCH=wasm go build -o assets/core.wasm cmd/main.go