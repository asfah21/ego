package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes mendaftarkan semua route ke Gin engine
func SetupRoutes(r *gin.Engine) {
	// 1. Halaman Utama
	r.GET("/", ShowHome)

	// 2. Halaman Kuesioner
	r.GET("/quiz", ShowQuiz)

	// 3. Proses Jawaban (HTMX)
	r.POST("/submit-tes", SubmitTest)

	// 4. Paywall
	r.GET("/paywall/:id", ShowPaywall)

	// 5. Hasil Premium (hanya jika PAID)
	r.GET("/hasil/:id", ShowResult)

	// 6. Halaman Informasi
	r.GET("/testimoni", ShowTestimoni)
	r.GET("/faq", ShowFAQ)
	r.GET("/produk", ShowProduk)

	// 7. Handle 404
	r.NoRoute(Show404)

	// 8. Recovery Middleware untuk 5xx
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.HTML(http.StatusInternalServerError, "error_5xx.html", nil)
				c.Abort()
			}
		}()
		c.Next()
	})
}
