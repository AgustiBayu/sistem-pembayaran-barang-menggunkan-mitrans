# SISTEM PEMBAYARAN BARANG MENGGUNKAN MIDTRANS
Sistem Pembayaran Barang Menggunakan Midtrans adalah sebuah solusi pembayaran online yang memungkinkan pengguna untuk melakukan transaksi pembelian barang secara aman dan efisien melalui platform Midtrans. 

![GitHub Logo](https://cdn.prod.website-files.com/6100d0111a4ed76bc1b9fd54/62217e885f52b860da9f00cc_Apa%20Itu%20Golang%3F%20Apa%20Saja%20Fungsi%20Dan%20Keunggulannya%20-%20Binar%20Academy.jpeg)

## Fitur Utama
- **Integrasi Payment Gateway:** Sistem ini mengintegrasikan API Midtrans untuk memproses transaksi pembayaran barang.   
- **Manajemen Produk:** Terdapat proses implementasi CRUD pada produk dalam pengolahannya.
- **Manajemen Pesanan:** Terdapat proses implementasi CRUD pada pesanan dalam pengolahannya.
- **Manajemen Pelanggan:** Terdapat proses implementasi CRUD pada pelanggan dalam pengolahannya.
- **Manajemen Order Item:** Terdapat proses implementasi CRUD pada order item dalam pengolahannya.
- **Manajemen Payment:** Terdapat proses implementasi CRUD pada payment dalam pengolahannya.
- **Continuous Integration:** GitHub Actions digunakan untuk otomatisasi proses build dan pengecekan kode. Setiap kali kode di-push ke branch `master`, workflow akan berjalan untuk memastikan bahwa aplikasi dibangun tanpa error.

## Teknologi
- **Bahasa:** Golang
- **Midtrans Go SDK:**
    ```bach
    github.com/midtrans/midtrans-go
- **Golang Migrate:**
    ```bach
    go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- **Golang Httprouter:**
    ```bach
    github.com/julienschmidt/httprouter
- **Golang Validate:**
    ```bach
    github.com/go-playground/validator/v10
- **Database:** PostgreSQL
    ```bach
    github.com/lib/pq
## Instalasi
1. Clone repository:
   ```bash
   git clone https://github.com/AgustiBayu/sistem-pembayaran-barang-menggunkan-mitrans.git
   
2. cd sistem-manajemen-restoran
3. go mod tidy
4. Atur konfigurasi database di file app.
5. Jalankan Perintah
   ```bash
   migrate -database "postgres://postgres:password!!@localhost:5432/db_name?sslmode=disable" -path migrations up
6. Jalankan aplikasi:
    ```bash
    go run main.go
