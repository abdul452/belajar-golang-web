# Belajar Golang Web

Repositori ini berisi kumpulan modul, latihan kode, dan unit test yang saya bangun selama mempelajari pemrograman backend web menggunakan bahasa **Go (Golang)**. Seluruh materi di dalam proyek ini diimplementasikan dengan memanfaatkan **Standard Library Go** (`net/http`) tanpa ketergantungan pada framework pihak ketiga (*zero-dependency framework approach*), guna memahami fondasi arsitektur web secara mendalam.

## 🚀 Fitur & Topik Pembelajaran

Proyek ini dibagi menjadi beberapa modul terisolasi berdasarkan silabus kompetensi backend:

1. **Server & Handler Basikal** (`01` - `03`)
   * Mengonfigurasi dan menjalankan HTTP Server menggunakan `http.Server`.
   * Memahami implementasi `http.Handler` dan `http.HandlerFunc` untuk memproses *request*.
2. **Routing & Request Data** (`04` - `08`)
   * Manajemen routing menggunakan `http.NewServeMux`.
   * Ekstraksi URL Query Parameter dan manipulasi HTTP Headers.
   * Pemrosesan data form melalui metode `POST` (`Form Post`).
   * Manajemen status menggunakan HTTP Cookies.
3. **Static Files & File Upload** (`09`, `10`, `18`, `19`)
   * Melayani aset statis (CSS, JS, Gambar) menggunakan `http.FileServer`.
   * Implementasi fitur unggah berkas (*File Upload*) dengan enkapsulasi `multipart/form-data`.
   * Mekanisme unduh berkas otomatis (*File Download*) memanfaatkan query parameter.
4. **Go HTML Template Engine** (`11` - `16`)
   * Rendering halaman web dinamis berbasis data `Struct` dan `Map`.
   * Implementasi *Template Actions* (Kondisional, Perulangan, Pipilines, dan Global Functions).
   * Fitur keamanan internal *Auto-Escaping* terhadap celah kerentanan XSS (Cross-Site Scripting).
   * Optimalisasi performa server menggunakan *Template Caching* berbasis `embed.FS`.
5. **Advanced Web Context** (`17`, `20`)
   * Mekanisme HTTP Redirects (Temporary & Permanent).
   * Arsitektur **Middleware** (Interceptors) untuk penanganan otentikasi, logging, dan pemulihan *panic error* secara terpusat.

---

## 📂 Struktur Repositori

```text
.
├── 01-server-test/
├── 02-handler-test/
├── 03-http-test/
├── 04-query-param-test/
├── 05-header-test/
├── 06-form-post-test/
├── 07-response-code-test/
├── 08-cookie-test/
├── 09-file-server-test/
│   └── resources/          # Aset statis (index.html, css, js)
├── 10-serve-file-test/
├── 11-template-test/
├── 12-template-data-test/
├── 13-template-action-test/
├── 14-template-layout-test/
├── 15-template-function-test/
├── 16-template-caching-test/
├── 17-redirect_test/
├── 18-upload-file-test/
│   ├── resources/          # Direktori penyimpanan hasil upload
│   └── templates/          # Berkas .gohtml untuk UI Upload
├── 19-download-file-test/
├── 20-middleware-test/
├── go.mod
└── README.md
```
## 🛠️ Prasyarat Sistem
- Go Compiler: Versi `1.20` atau yang lebih baru.
- Sistem Operasi: Linux (Ubuntu / Termux), macOS, atau Windows.

## 🧪 Cara Menjalankan Unit Test
Setiap modul di dalam repositori ini dibangun menggunakan pendekatan Test-Driven Development (TDD) mini menggunakan package bawaan `testing` dan `net/http/httptest`.

Menjalankan Seluruh Pengujian
Untuk memverifikasi semua modul berfungsi dengan normal tanpa kendala, jalankan perintah berikut pada direktori utama:
```Bash
go test -v ./...
```
### Menjalankan Pengujian Spesifik Per Modul
Masuk ke direktori modul yang diinginkan, kemudian jalankan test-nya:
```Bash
cd 18-upload-file-test
go test -v -run ^TestUploadForm$ -timeout 0
```
Catatan: Gunakan parameter `-timeout 0` pada test yang menginisiasi server HTTP aktif (`ListenAndServe`) agar proses pengujian tidak dihentikan paksa setelah 30 detik.

## 💡 Konsep Utama yang Dikuasai
- Concurrency-Ready: Memahami bagaimana Go melayani setiap HTTP request yang masuk di dalam Goroutine yang terpisah secara efisien.

- Security Awareness: Memahami perbedaan krusial antara package `text/template` (tanpa proteksi) dan `html/template` yang memiliki fitur pengamanan ketat terhadap manipulasi injeksi skrip berbahaya.

- Separation of Concerns (SoC): Mampu memisahkan data biner multipart dengan data teks parameter biasa (`request.FormValue`) saat menangani pengiriman data kompleks dari klien.
