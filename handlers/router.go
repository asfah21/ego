package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mendaftarkan semua route ke Gin engine
func SetupRoutes(r *gin.Engine) {
	// 0. Recovery Middleware untuk 5xx — harus PALING ATAS
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.HTML(http.StatusInternalServerError, "error_5xx.html", nil)
				c.Abort()
			}
		}()
		c.Next()
	})

	// 1. Halaman Utama (HEAD juga penting untuk curl -I dan health check OpenResty)
	r.GET("/", ShowHome)
	r.HEAD("/", ShowHome)

	// 2. Halaman Kuesioner
	r.GET("/quiz", ShowQuiz)
	r.HEAD("/quiz", ShowQuiz)

	// 3. Proses Jawaban (HTMX)
	r.POST("/submit-tes", SubmitTest)

	// 4. Paywall
	r.GET("/paywall/:id", ShowPaywall)
	r.HEAD("/paywall/:id", ShowPaywall)

	// 5. Hasil Premium (hanya jika PAID)
	r.GET("/hasil/:id", ShowResult)
	r.HEAD("/hasil/:id", ShowResult)

	// 6. Halaman Informasi
	r.GET("/testimoni", ShowTestimoni)
	r.HEAD("/testimoni", ShowTestimoni)
	r.GET("/faq", ShowFAQ)
	r.HEAD("/faq", ShowFAQ)
	r.GET("/produk", ShowProduk)
	r.HEAD("/produk", ShowProduk)
	r.GET("/tentang", ShowTentang)
	r.HEAD("/tentang", ShowTentang)

	// 7. Handle 404 — harus di PALING AKHIR
	r.NoRoute(Show404)
}
