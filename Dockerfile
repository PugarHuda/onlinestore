# Menggunakan golang sebagai base image
FROM golang:alpine AS builder

# Menambahkan dependensi yang diperlukan
RUN apk update && apk add --no-cache git

# Menentukan direktori kerja di dalam kontainer
WORKDIR /app

# Menyalin file Go mod dan sum untuk mengambil dependensi
COPY go.mod ./
COPY go.sum ./

# Men-download dan menginstal dependensi
RUN go mod download

# Menyalin kode aplikasi ke dalam kontainer
COPY . .

# Membangun aplikasi Go
RUN go build -o main .

# Tahap kedua, membuat citra akhir yang lebih kecil
FROM alpine:latest  

# Menentukan direktori kerja di dalam kontainer
WORKDIR /root/

# Menyalin executable yang telah dibangun dari tahap sebelumnya
COPY --from=builder /app/main .

# Menjalankan aplikasi ketika kontainer dimulai
CMD ["./main"]
