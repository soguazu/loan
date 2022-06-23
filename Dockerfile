FROM 1.18.3-alpine3.16 as builder

RUN apk --no-cache add gcc g++ make git

# Define build env

ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app

COPY . .
# Build app

RUN go mod init core_business

RUN go mod tidy

RUN GOOS=linux go build -ldflags="-s -w" -o app ./cmd/main.go

FROM alpine:3.15 as production

# Copy built binary from builder
COPY --from=builder app .

# Expose port
EXPOSE 30300
# Exec built binary
#CMD ["./app"]

ENTRYPOINT /app --port 30300


