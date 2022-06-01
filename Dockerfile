# syntax=docker/dockerfile:1

# STEP 1 create user
FROM alpine:3.15 as authority
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

# STEP 2 build executable binary
FROM golang:1.18-alpine AS builder
WORKDIR /src
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -tags nomsgpack \
    -o /app ./cmd/restapp

# STEP 3 build a scratch image
FROM scratch
LABEL maintainer="marcosvieirajr@gmail.com"
WORKDIR /
COPY --from=authority /etc/passwd /etc/passwd
COPY --from=authority /etc/group /etc/group
COPY --from=builder /app .
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=user
ENV DB_PASSWORD=secret
ENV DB_NAME=dbname
EXPOSE 3000
USER appuser:appuser
ENTRYPOINT ["/app"]