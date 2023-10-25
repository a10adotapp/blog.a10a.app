FROM golang:1.21 AS build

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o /go/bin/app /go/src/app

# Now copy it into our base image.
FROM scratch AS runner

COPY --from=build /go/bin/app /

CMD ["/app"]

EXPOSE 3000
