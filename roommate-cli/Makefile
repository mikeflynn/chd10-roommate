all: deps format data
		@echo "--> Building"
		@go build -o ../ComputerRoommate/roommate-cli

deps:
		@echo "--> Installing build dependencies"
		@go get github.com/deckarep/gosx-notifier
		@go get github.com/jteeuwen/go-bindata

format:
		@echo "--> Running go fmt"
		@gofmt -s -w .

data:
		@echo "--> Importing binary files"
		@go-bindata -o scripts.go scripts/