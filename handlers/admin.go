package handlers

import (
	"net/http"

	"ego/services"

	"github.com/gin-gonic/gin"
)

// Admin credentials (hardcoded)
const (
	adminUsername = "admin"
	adminPassword = "admin360"
)

// ShowLogin menampilkan halaman login admin
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// LoginProcess memproses login admin
func LoginProcess(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != adminUsername || password != adminPassword {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Error": "Username atau password salah",
		})
		return
	}

	// Set session cookie (simple token-based)
	c.SetCookie("admin_token", "authenticated", 3600*2, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

// ShowDashboard menampilkan dashboard admin dengan data user
func ShowDashboard(c *gin.Context) {
	// Cek autentikasi
	token, err := c.Cookie("admin_token")
	if err != nil || token != "authenticated" {
		c.Redirect(http.StatusSeeOther, "/admin/login")
		return
	}

	users, err := services.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error_5xx.html", nil)
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Users": users,
	})
}

// LogoutProcess menghapus session admin
func LogoutProcess(c *gin.Context) {
	c.SetCookie("admin_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/login")
}
