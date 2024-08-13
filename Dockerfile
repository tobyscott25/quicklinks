FROM golang:1.22.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests into the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GIN_MODE=release
# ENV GOOS=linux GOARCH=amd64
RUN go build -o /bin/quick-links .


FROM alpine:3.14

RUN addgroup qluser && adduser -D -G qluser qluser

# Copy the built Go binary from the builder stage
COPY --chown=qluser:qluser --from=builder /bin/quick-links /bin/quick-links

USER qluser:qluser

EXPOSE 8080

CMD ["/bin/quick-links"]