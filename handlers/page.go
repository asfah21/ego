package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowHome menampilkan halaman utama
func ShowHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowQuiz menampilkan halaman kuesioner
func ShowQuiz(c *gin.Context) {
	c.HTML(http.StatusOK, "quiz.html", nil)
}

// ShowTestimoni menampilkan halaman testimoni
func ShowTestimoni(c *gin.Context) {
	c.HTML(http.StatusOK, "testimoni_page.html", nil)
}

// ShowFAQ menampilkan halaman FAQ
func ShowFAQ(c *gin.Context) {
	c.HTML(http.StatusOK, "faq_page.html", nil)
}

// ShowProduk menampilkan halaman produk
func ShowProduk(c *gin.Context) {
	c.HTML(http.StatusOK, "produk_page.html", nil)
}

// ShowTentang menampilkan halaman tentang kami
func ShowTentang(c *gin.Context) {
	c.HTML(http.StatusOK, "tentang.html", nil)
}

// Show404 menampilkan halaman 404
func Show404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error_404.html", nil)
}

// Show500 menampilkan halaman error server
func Show500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "error_5xx.html", nil)
}
