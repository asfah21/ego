package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Inisialisasi Database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL env is missing")
	}

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Database init error: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	log.Println("⚡ Terkoneksi ke PostgreSQL dengan sukses!")

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Load HTML Templates modular (urutan penting):
	// 1. sections/ → define komponen (navbar, footer, produk, testimoni, faq, error pages)
	// 2. quiz/ → define pertanyaan (q1, q2, q3)
	// 3. layout.html → base template dengan head, navbar, footer
	// 4. _pages/* → halaman yang memanggil {{template "layout"}}
	// 5. index.html → halaman utama (tidak pakai layout, langsung sections)
	templ := template.Must(template.ParseGlob("templates/sections/*.html"))
	templ = template.Must(templ.ParseGlob("templates/quiz/*.html"))
	templ = template.Must(templ.ParseGlob("templates/layout.html"))
	templ = template.Must(templ.ParseGlob("templates/_pages/*.html"))
	templ = template.Must(templ.ParseGlob("templates/index.html"))
	r.SetHTMLTemplate(templ)
	r.Static("/assets", "./assets")

	// Custom Recovery Middleware untuk 5xx — tampilkan halaman error yang cantik
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.HTML(http.StatusInternalServerError, "error_5xx.html", nil)
				c.Abort()
			}
		}()
		c.Next()
	})

	// 1. Route Halaman Utama
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Route Halaman Kuesioner Terpisah
	r.GET("/quiz", func(c *gin.Context) {
		c.HTML(http.StatusOK, "quiz.html", nil)
	})

	// 2. Route Proses Jawaban (Ditembak menggunakan HTMX)
	r.POST("/submit-tes", processTest)

	// 3. Route Paywall (Fake Loading AI + Tombol Bayar)
	r.GET("/paywall/:id", showPaywall)

	// 4. Route Hasil Premium (Hanya terbuka jika status PAID)
	r.GET("/hasil/:id", showResult)

	// 5. Route Halaman Informasi Modular
	r.GET("/testimoni", func(c *gin.Context) {
		c.HTML(http.StatusOK, "testimoni_page.html", nil)
	})
	r.GET("/faq", func(c *gin.Context) {
		c.HTML(http.StatusOK, "faq_page.html", nil)
	})
	r.GET("/produk", func(c *gin.Context) {
		c.HTML(http.StatusOK, "produk_page.html", nil)
	})

	// 6. Handle 404 - Route tidak ditemukan
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error_404.html", nil)
	})

	// Jalankan di port 8080 (sesuai Dockerfile)
	r.Run(":8080")
}

// Fungsi memproses skor dari form HTML
func processTest(c *gin.Context) {
	nama := c.PostForm("nama")
	email := c.PostForm("email")

	// Kita punya 3 pertanyaan untuk contoh (bisa ditambah nanti)
	// Skor 1-5 (Sangat Tidak Setuju -> Sangat Setuju)
	q1, _ := strconv.Atoi(c.PostForm("q1")) // Narsisme
	q2, _ := strconv.Atoi(c.PostForm("q2")) // Machiavellian
	q3, _ := strconv.Atoi(c.PostForm("q3")) // Psikopati

	// Hitung total skor (Logika di backend aman dari manipulasi browser)
	skorNarsisme := q1 * 20 // Jadikan persentase maks 100
	skorMachiavellian := q2 * 20
	skorPsikopati := q3 * 20

	// Simpan ke Postgres dan dapatkan UUID user tersebut
	var userID string
	query := `INSERT INTO users_test (nama, email, skor_narsisme, skor_machiavellian, skor_psikopati) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(query, nama, email, skorNarsisme, skorMachiavellian, skorPsikopati).Scan(&userID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan data tes: "+err.Error())
		return
	}

	// 💡 TRIK HTMX: Alihkan user ke halaman paywall menggunakan header HX-Redirect
	c.Header("HX-Redirect", fmt.Sprintf("/paywall/%s", userID))
	c.Status(http.StatusOK)
}

func showPaywall(c *gin.Context) {
	id := c.Param("id")
	// Ambil data user dari DB untuk personalisasi halaman paywall
	var nama string
	err := db.QueryRow("SELECT nama FROM users_test WHERE id = $1", id).Scan(&nama)
	if err != nil {
		c.String(http.StatusNotFound, "User tidak ditemukan")
		return
	}

	c.HTML(http.StatusOK, "paywall.html", gin.H{
		"ID":   id,
		"Nama": nama,
	})
}

func showResult(c *gin.Context) {
	id := c.Param("id")
	var nama, status string
	var narsisme, mach, psikopat int

	query := "SELECT nama, skor_narsisme, skor_machiavellian, skor_psikopati, status_pembayaran FROM users_test WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&nama, &narsisme, &mach, &psikopat, &status)
	if err != nil {
		c.String(http.StatusNotFound, "Data tidak ditemukan")
		return
	}

	// Proteksi Paywall: Kalau belum bayar, tendang balik ke halaman paywall!
	if status != "PAID" {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/paywall/%s?error=belum_bayar", id))
		return
	}

	c.HTML(http.StatusOK, "hasil.html", gin.H{
		"Nama":          nama,
		"Narsisme":      narsisme,
		"Machiavellian": mach,
		"Psikopati":     psikopat,
	})
}
