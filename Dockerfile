# ---- build stage ----
FROM golang:1.22 as build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/myservice ./cmd/myservice

# ---- run stage ----
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/myservice /app/myservice
ENV PORT=8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app/myservice"]