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

	// 1. Parse sections (navbar, footer, produk_section, etc.)
	//    These define named templates like "navbar", "footer", "produk_section"
	all := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/sections/*.html"))

	// 2. Parse layout — this defines the outer HTML structure
	//    layout.html has NO {{define}}, so it becomes template "layout.html"
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/layout.html"))

	// 3. Parse index.html first — it defines {{define "content"}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/index.html"))

	// 4. Parse all page files — each defines {{define "content"}}
	//    They all define the SAME template name "content", so the LAST one parsed wins.
	//    This is fine because we only use one page at a time.
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/_pages/*.html"))

	// 5. Create wrapper templates named after each page file so Gin's c.HTML() can find them.
	//    Each wrapper simply calls {{template "layout.html" .}} which renders the full page.
	//    layout.html calls {{template "content" .}} which uses the "content" template from the page.
	pageFiles := []string{
		"index.html",
		"quiz.html",
		"paywall.html",
		"hasil.html",
		"tentang.html",
		"login.html",
		"dashboard.html",
		"user_detail.html",
		"error.html",
	}

	for _, name := range pageFiles {
		// Create a template named e.g. "index.html" that renders the layout
		// The layout calls {{template "content" .}} which is defined in the page file
		template.Must(all.New(name).Parse(`{{template "layout.html" .}}`))
	}

	r.SetHTMLTemplate(all)
	r.Static("/assets", "./assets")
}
