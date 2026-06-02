package templates

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// Setup memuat semua template HTML dan static assets
func Setup(r *gin.Engine) {
	templ := template.Must(template.ParseGlob("templates/sections/*.html"))
	templ = template.Must(templ.ParseGlob("templates/quiz/*.html"))
	templ = template.Must(templ.ParseGlob("templates/layout.html"))
	templ = template.Must(templ.ParseGlob("templates/_pages/*.html"))
	templ = template.Must(templ.ParseGlob("templates/index.html"))
	r.SetHTMLTemplate(templ)
	r.Static("/assets", "./assets")
}
