do:
	go build -o release/b fun.go
	GOOS=windows go build -o release/b.exe fun.go
	GOOS=android GOARCH=arm64 go build -o release/a fun.go
