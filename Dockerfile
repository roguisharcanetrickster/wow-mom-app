FROM golang:1.23-alpine AS build

WORKDIR /app

# Install build dependencies for CGO/SQLite
RUN apk add --no-cache gcc musl-dev

# Move into the subfolder where the Go code lives
WORKDIR /app/wow-mom-app

COPY wow-mom-app/go.mod wow-mom-app/go.sum ./
RUN go mod download

# Copy the entire workspace to get the subfolder content
WORKDIR /app
COPY . .

# Build from the subfolder
WORKDIR /app/wow-mom-app
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/server main.go

FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates sqlite-libs

COPY --from=build /app/server /app/server
# Copy static/templates from the subfolder
COPY --from=build /app/wow-mom-app/static /app/wow-mom-app/static
COPY --from=build /app/wow-mom-app/templates /app/wow-mom-app/templates

# Fix path expectations: the app expects them in ./static and ./templates relative to binary or WD
RUN mv /app/wow-mom-app/static /app/static && mv /app/wow-mom-app/templates /app/templates

EXPOSE 8080

CMD ["/app/server"]
