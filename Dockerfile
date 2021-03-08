FROM golang:1.15

WORKDIR /app
RUN apt-get update && apt-get install git -y

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

FROM env as build
COPY . .
RUN make build

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/catalog /bin/catalog
ENTRYPOINT ["/bin/catalog"]
