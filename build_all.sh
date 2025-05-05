GOOS=darwin go build -o xmltomf
GOOS=linux go build -o xmltomf_linux
GOOS=windows GOARCH=386 go build -o xmltomf.exe