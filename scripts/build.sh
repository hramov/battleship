export PATH=$PATH:/usr/local/go/bin
echo "Building the app"
go build -o ./bin/main ./cmd/main.go
echo "Successfully built the app"
