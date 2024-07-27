# syntax=docker/dockerfile:1

FROM golang:1.22.4

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /reverse-uber-be
EXPOSE 8080

# Run
CMD ["/reverse-uber-be"]
