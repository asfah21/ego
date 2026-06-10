package templates

import (
	"html/template"
	"math"
	"unicode"
	"unicode/utf8"

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
		"firstChar": func(s string) string {
			if s == "" {
				return "?"
			}
			firstRune, _ := utf8.DecodeRuneInString(s)
			return string(unicode.ToUpper(firstRune))
		},
	}

	templ := template.New("").Funcs(funcMap)

	// Load sections (navbar, footer, etc.)
	templ = template.Must(templ.ParseGlob("templates/sections/*.html"))

	// Load layout
	templ = template.Must(templ.ParseGlob("templates/layout.html"))

	// Load pages
	templ = template.Must(templ.ParseGlob("templates/_pages/*.html"))

	// Load index (home page with embedded sections)
	templ = template.Must(templ.ParseGlob("templates/index.html"))

	r.SetHTMLTemplate(templ)
	r.Static("/assets", "./assets")
}
