git add .
git commit -m "New commit"
git push

$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o bootstrap main.go

Remove-Item function.zip
Compress-Archive -Path .\bootstrap -DestinationPath function.zip