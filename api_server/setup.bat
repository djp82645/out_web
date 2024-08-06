set GOARCH=amd64
set GOOS=linux
set BuildVersion=1.0.0
gofumpt -l -w ./
go build -ldflags="-X main.BuildVersion=%BuildVersion% -s -w" -o api_server main.go && upx -9 api_server 

