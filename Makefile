wasm:
	GOOS=js GOARCH=wasm go build -o assets/go.wasm cmd/main.go