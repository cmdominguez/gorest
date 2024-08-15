# Establecer la imagen base de Go
FROM golang:1.22-alpine

# Instalar git y herramientas necesarias
RUN apk add --no-cache git

# Instalar 'air' para recarga en caliente
RUN go install github.com/air-verse/air@latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum para instalar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Exponer el puerto que utilizará Gin
EXPOSE 8080

# Comando para ejecutar la aplicación usando 'air'
CMD ["air", "-c", ".air.toml"]