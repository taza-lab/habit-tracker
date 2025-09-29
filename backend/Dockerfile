# ビルドステージ
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 依存関係の解決
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードのコピーとビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/habit-tracker

# 実行ステージ
FROM alpine:latest

WORKDIR /root/

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .

# ポートを公開
EXPOSE 8080

# アプリケーションの実行
CMD ["./main"]
