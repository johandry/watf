default: build serve

build:
	GOARCH=wasm GOOS=js go build -o lib.wasm main.go

serve:
	go run server.go