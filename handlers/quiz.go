package handlers

import (
	"net/http"
	"strconv"

	"ego/services"

	"github.com/gin-gonic/gin"
)

// SubmitTest memproses jawaban kuis dari form HTMX
func SubmitTest(c *gin.Context) {
	nama := c.PostForm("nama")
	email := c.PostForm("email")

	q1, _ := strconv.Atoi(c.PostForm("q1"))
	q2, _ := strconv.Atoi(c.PostForm("q2"))
	q3, _ := strconv.Atoi(c.PostForm("q3"))

	userID, err := services.ProcessQuizAnswers(nama, email, q1, q2, q3)
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal menyimpan data tes: "+err.Error())
		return
	}

	// HTMX redirect ke halaman paywall
	c.Header("HX-Redirect", "/paywall/"+userID)
	c.Status(http.StatusOK)
}

// ShowPaywall menampilkan halaman pembayaran
func ShowPaywall(c *gin.Context) {
	id := c.Param("id")
	data, err := services.GetPaywallData(id)
	if err != nil {
		c.String(http.StatusNotFound, "User tidak ditemukan")
		return
	}

	c.HTML(http.StatusOK, "paywall.html", gin.H{
		"ID":   data.ID,
		"Nama": data.Nama,
	})
}

// ShowResult menampilkan hasil kuis (hanya jika sudah bayar)
func ShowResult(c *gin.Context) {
	id := c.Param("id")
	result, err := services.GetQuizResult(id)
	if err != nil {
		c.String(http.StatusNotFound, "Data tidak ditemukan")
		return
	}
	if result == nil {
		// Belum bayar, redirect ke paywall
		c.Redirect(http.StatusSeeOther, "/paywall/"+id+"?error=belum_bayar")
		return
	}

	c.HTML(http.StatusOK, "hasil.html", gin.H{
		"Nama":          result.Nama,
		"Narsisme":      result.Narsisme,
		"Machiavellian": result.Machiavellian,
		"Psikopati":     result.Psikopati,
	})
}
