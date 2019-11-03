all: visp.exe

visp.exe: main.go
	GOOS=windows GOARCH=amd64 go build -o visp.exe -ldflags="-H windowsgui" main.go

rm:
	rm visp.exe
