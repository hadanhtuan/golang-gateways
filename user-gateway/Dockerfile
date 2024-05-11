FROM golang:1.21-alpine as build

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /user-gateway /user-gateway/cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /go/src/app

COPY --from=build /user-gateway /user-gateway

EXPOSE 3000

CMD ["/user-gateway"]