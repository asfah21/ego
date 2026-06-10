package templates

import (
	"html/template"
	"math"

	"github.com/gin-gonic/gin"
)

// Setup memuat semua template HTML dan static assets
func Setup(r *gin.Engine) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"round": func(f float64) float64 {
			return math.Round(f)
		},
	}

	templ := template.New("").Funcs(funcMap)
	templ = template.Must(templ.ParseGlob("templates/sections/*.html"))
	templ = template.Must(templ.ParseGlob("templates/layout.html"))
	templ = template.Must(templ.ParseGlob("templates/_pages/*.html"))
	templ = template.Must(templ.ParseGlob("templates/index.html"))
	r.SetHTMLTemplate(templ)
	r.Static("/assets", "./assets")
}
