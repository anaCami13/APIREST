FROM golang:1.23.2

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app 

COPY api-res/go.* ./

RUN go mod download

COPY api-res/. .

# Copia el script wait-for-it.sh al contenedor y colócalo en /usr/local/bin
COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh

# Otorga permisos de ejecución al script
RUN chmod +x /usr/local/bin/wait-for-it.sh

# Construye la aplicación y la coloca en /usr/local/bin
RUN go build -v -o /usr/local/bin/app .

EXPOSE 3000

# Usa la ruta completa para ejecutar wait-for-it.sh
CMD ["/usr/local/bin/wait-for-it.sh", "mariadb:3306", "--", "/usr/local/bin/app"]



