# Belajar Golang dengan Gin

Projek ini adalah contoh sederhana untuk belajar bahasa pemrograman Go (Golang) menggunakan framework web Gin untuk penanganan rute (routing). Projek ini memiliki beberapa contoh dasar tentang bagaimana menggunakan Gin untuk membuat aplikasi web sederhana.

## Instalasi

Pastikan Anda telah menginstal Go di komputer Anda. Untuk memulai, Anda juga perlu menginstal framework Gin dengan menjalankan perintah berikut di terminal:

```
go get -u github.com/gin-gonic/gin
```

Setelah menginstal Gin, Anda dapat menjalankan projek ini dengan menyalin repositori ke komputer Anda dan menjalankan langkah-langkah berikut di terminal:

1. Build Project dengan menjalankan peringah :

```
go build
```

2. Jika proses build telah selesai anda dapat menjalankan file hasil build yang bernama `GO-learn_gin.exe`. Ini akan memulai server web lokal, dan Anda dapat mengakses aplikasi melalui browser dengan alamat `http://localhost:8080`.

## Struktur Projek

- `main.go`: Berkas utama yang berisi logika untuk menginisialisasi server dan menentukan rute.
- `handlers.go`: Berkas ini berisi definisi penanganan permintaan (request handlers) untuk setiap rute yang ditentukan dalam aplikasi.
- `models.go`: Berkas ini berisi data-data artikel yang ada pada aplikasi ini.
- `README.md`: Berkas ini! Berisi dokumentasi untuk menjelaskan tentang projek ini dan cara menjalankannya.

## Contoh Penggunaan

Projek ini memiliki beberapa contoh rute sederhana yang dapat Anda gunakan untuk memahami penggunaan dasar Gin. Berikut adalah contoh beberapa rute yang telah disediakan:

- `http://localhost:8080`: Menampilkan semua artikel yang ada pada aplikasi ini.
- `http://localhost:8080/article/view/:article_id`: Menampilkan detail data suati artikel dengan `article_id` sebagai identifikasi data.

Anda dapat menambahkan atau mengubah penanganan rute sesuai kebutuhan projek Anda.
