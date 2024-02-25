FROM golang:1.22.0 AS build
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM scratch
COPY --from=build /work/main /
ENTRYPOINT ["/main"]
