FROM coopsafe/pdf-generator-base:latest
WORKDIR /app
COPY . .
RUN go mod tidy
RUN	go mod download
ENV TZ America/Sao_Paulo
RUN GOBIN=/usr/local/bin/ GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags="-s -w" ./cmd/pdf-generation
EXPOSE 8080
CMD pdf-generation