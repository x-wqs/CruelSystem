# MIT 6.824

## Build
- ./.check-build

## Lab 1 - Sequential
- cd src/main
- go build -race -buildmode=plugin ../mrapps/wc.go
- rm mr-out*
- go run -race mrsequential.go wc.so pg*.txt
- more mr-out-0




