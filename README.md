# 📂 Project App Portfolio (Golang)

![Go Version](https://img.shields.io/badge/Go-1.25.5-blue?style=flat&logo=go)
![Database](https://img.shields.io/badge/PostgreSQL-16-blue?style=flat&logo=postgresql)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20Layered-green?style=flat)

Aplikasi web portfolio pribadi yang dibangun menggunakan bahasa pemrograman **Golang**. Proyek ini menerapkan konsep **Clean Architecture** (Handler, Service, Repository), rendering HTML Template server-side, serta implementasi **Structured Logging** yang robust.

---

## 📋 Daftar Isi
- [Fitur Utama](#-fitur-utama)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Struktur Proyek](#-struktur-proyek)
- [Prasyarat](#-prasyarat)
- [Instalasi dan Cara Menjalankan](#-instalasi-dan-cara-menjalankan)
- [Konfigurasi Database](#-konfigurasi-database)
- [Logging System](#-logging-system)

---

## ✨ Fitur Utama
1.  **Halaman Web Dinamis**: Menggunakan `html/template` untuk merender halaman Home, About, Portfolio, Resume, dan Contact.
2.  **Clean Architecture**: Kode terorganisir dengan pemisahan tugas yang jelas antara:
    -   **Handler**: Menangani HTTP Request/Response.
    -   **Service**: Berisi logika bisnis.
    -   **Repository**: Menangani interaksi langsung ke database.
3.  **Database PostgreSQL**: Penyimpanan data persisten untuk projects, skills, experience, dan pesan masuk.
4.  **Structured Logging**: Menggunakan **Uber Zap** & **Lumberjack** untuk mencatat aktivitas aplikasi dan error ke file log dengan fitur rotasi otomatis.
5.  **Middleware Logging**: Mencatat setiap request HTTP (Method, Path, Duration) yang masuk ke server.
6.  **Form Kontak**: Pengunjung dapat mengirim pesan yang akan tersimpan langsung ke database.

---

## 🛠 Teknologi yang Digunakan

| Kategori | Teknologi | Deskripsi |
| :--- | :--- | :--- |
| **Backend** | Go (Golang) v1.25.5 | Bahasa pemrograman utama |
| **Database** | PostgreSQL | Relational Database Management System |
| **Router** | [go-chi/chi](https://github.com/go-chi/chi) v5 | Router HTTP yang ringan dan cepat |
| **DB Driver** | [pgx](https://github.com/jackc/pgx) v5 | Driver PostgreSQL performance tinggi |
| **Logger** | [uber-go/zap](https://github.com/uber-go/zap) | Structured Logging |
| **Frontend** | HTML5, CSS3, JS | Tampilan antarmuka pengguna |

---

## 📂 Struktur Proyek

```text
.
├── assets/              # File statis (CSS, JS, Images)
├── database/            # Konfigurasi koneksi database
├── handler/             # Layer Handler (Controller/Delivery)
├── logs/                # Output file log aplikasi (Auto-generated)
├── middleware/          # Middleware HTTP (Logging)
├── model/               # Definisi Struct data (Entities)
├── repository/          # Layer akses data (Query SQL)
├── service/             # Layer logika bisnis (Use Case)
├── templates/           # File HTML Templates
├── utils/               # Fungsi utilitas (Init Logger, Response)
├── main.go              # Entry point aplikasi
├── go.mod               # Dependency management
├── portofolio           # File Dump Database SQL
└── README.md            # Dokumentasi proyek
```

##Video Tutorial : https://drive.google.com/file/d/1_P6xsKeHENW08PCc2irqUx6HrgkfTuCV/view?usp=sharing
