all:
	@@go build image_analyzer.go
linux:
	@@GOOS=linux GOARCH=386 go build image_analyzer.go
