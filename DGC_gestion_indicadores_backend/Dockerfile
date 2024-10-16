FROM golang:1.22.0

WORKDIR /dgc_gestion_indicadores_backend

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /dgc_gestion_indicadores_backend

CMD ["/dgc_gestion_indicadores_backend"]