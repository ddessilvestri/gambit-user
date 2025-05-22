git add .
git commit -m "New commit"
git push

Set-Variable GOOS=linux
Set-Variable GOARCH=amd64

go build -o bootstrap main.go
Remove-Item function.zip
Compress-Archive -Path .\bootstrap -DestinationPath function.zip