APP_NAME=ec2control
VERSION=2.0.0

test:
	go test -v ./...
vet:
	go vet -v ./...
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux/$(APP_NAME) cmd/ec2/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows/$(APP_NAME).exe cmd/ec2/main.go
