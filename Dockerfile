FROM golang:1.23.2

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app 

COPY api-res/go.* ./

RUN go mod download

COPY api-res/. .

# Copia el script wait-for-it.sh al contenedor y otorga permisos de ejecución
COPY wait-for-it.sh ./

RUN go build -v -o /usr/local/bin/app .

EXPOSE 3000

# Usa wait-for-it.sh para esperar a que MariaDB esté listo antes de iniciar la aplicación
CMD ["wait-for-it.sh", "mariadb:3306", "--", "/usr/local/bin/app"]


