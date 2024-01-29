# 使用 Go 官方映像作為基礎映像
FROM golang:latest

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下載依賴
RUN go mod download

# 複製源代碼文件到工作目錄
COPY . .

# 構建應用程序
RUN go build -o main .

# 暴露端口
EXPOSE 8082

# 執行應用程序
CMD ["./main"]