git:
	git add .
	git commit -m "update"
	git push

windows:
	GOOS=windows GOARCH=386 go build -o sshclient.exe main.go

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o sshclient

macos:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o sshclient
