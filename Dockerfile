FROM golang:1.18 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o armo-webhook-server cmd/admission-controler/main.go

FROM gcr.io/distroless/base
COPY --from=build /app/armo-webhook-server /
EXPOSE 8443

CMD ["/armo-webhook-server"]
