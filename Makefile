build:
	go build -o bin/cpu-clamp main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=windows GOARCH=amd64 go build -o bin/clamp-windows.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/clamp-darwin main.go
	GOOS=linux GOARCH=amd64 go build -o bin/clamp main.go

all: hello build