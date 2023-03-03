FROM golang:1.18 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v ./...
RUN go test -v ./...

RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/layoffs/main.go

FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/app /
EXPOSE 3000
CMD ["/app"]