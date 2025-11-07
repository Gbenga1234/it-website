# Build stage
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod         --mount=type=cache,target=/root/.cache/go-build         go build -o ./server ./cmd/server

# Runtime stage (distroless-ish minimal)
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=build /app/server /app/server
COPY --from=build /app/web /app/web
ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
