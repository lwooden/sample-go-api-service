##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /sample-go-api-service

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /sample-go-api-service /sample-go-api-service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/sample-go-api-service"]