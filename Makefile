test:
	go test -v ./...
dev:
	go run ./cmd/main.go
build-linux:
	GOOS=linux GOARCH=amd64 go build -o fam-birthday ./cmd/main.go
build-windows:
	GOOS=windows GOARCH=amd64 go build -o fam-birthday.exe ./cmd/main.go
