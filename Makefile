test:
	go test -v ./...
test-coverage:
	go test -v -cover ./...
dev:
	go run ./cmd/main.go
build-linux-amd:
	GOOS=linux GOARCH=amd64 go build -o fam-birthday ./cmd/main.go
build-linux-arm:
	GOOS=linux GOARCH=arm64 go build -o fam-birthday ./cmd/main.go
build-windows:
	GOOS=windows GOARCH=amd64 go build -o fam-birthday.exe ./cmd/main.go
