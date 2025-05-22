git add .
git commit -m "New commit"
git push

# Configurar variables de entorno para compilar para Linux + x86_64
$env:GOOS = "linux"
$env:GOARCH = "amd64"

# Compilar el ejecutable Go con nombre 'bootstrap' (requerido por runtime provided.al2)
go build -o bootstrap main.go

# Eliminar archivo ZIP previo si existe
if (Test-Path "function.zip") {
    Remove-Item function.zip
}

# Crear un archivo ZIP con el ejecutable bootstrap
Compress-Archive -Path .\bootstrap -DestinationPath function.zip