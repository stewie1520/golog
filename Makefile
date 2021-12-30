compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--proto_path=. \
		--plugin=protoc-gen-go="/home/hieudong/go/bin/protoc-gen-go"
test:
	go test -race ./...

play:
	go run ./cmd/server/main.go