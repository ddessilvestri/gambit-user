git add .
git commit -m "last commit"
git push
set GOOS=linux
set GOARCH=amd64
go build main.go 
del main.zip
tar.exe -a -cf main.zip main 