all: visp.exe

visp.exe:
	GOOS=windows GOARCH=amd64 go build -o visp.exe main.go
