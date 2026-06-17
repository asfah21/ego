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
	all := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/sections/*.html"))

	// 2. Parse layout — defines outer HTML structure, calls {{template "content" .}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/layout.html"))

	// 3. Parse index.html — uses {{define "content"}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/index.html"))

	// 4. Parse each page file individually — each uses a unique define name
	//    e.g. {{define "content-quiz"}}, {{define "content-login"}}, etc.
	pageFiles := []string{
		"templates/_pages/quiz.html",
		"templates/_pages/paywall.html",
		"templates/_pages/hasil.html",
		"templates/_pages/tentang.html",
		"templates/_pages/login.html",
		"templates/_pages/dashboard.html",
		"templates/_pages/user_detail.html",
		"templates/_pages/error.html",
	}
	for _, f := range pageFiles {
		parsed := template.Must(template.Must(all.Clone()).ParseFiles(f))
		for _, t := range parsed.Templates() {
			if t.Name() != "" {
				all = template.Must(all.AddParseTree(t.Name(), t.Tree))
			}
		}
	}

	// 5. Create wrapper templates named after each page file so Gin's c.HTML() can find them.
	//    Each wrapper calls {{template "layout.html" .}} which renders the full page.
	//    layout.html calls {{template "content" .}} — so we must ensure "content" is defined.
	//    For index.html, "content" is defined directly in index.html.
	//    For other pages, we define "content" here in the wrapper to point to the unique content-* template.
	wrapperTemplates := map[string]string{
		"index.html":       `{{template "layout.html" .}}`,
		"quiz.html":        `{{define "content"}}{{template "content-quiz" .}}{{end}}{{template "layout.html" .}}`,
		"paywall.html":     `{{define "content"}}{{template "content-paywall" .}}{{end}}{{template "layout.html" .}}`,
		"hasil.html":       `{{define "content"}}{{template "content-hasil" .}}{{end}}{{template "layout.html" .}}`,
		"tentang.html":     `{{define "content"}}{{template "content-tentang" .}}{{end}}{{template "layout.html" .}}`,
		"login.html":       `{{define "content"}}{{template "content-login" .}}{{end}}{{template "layout.html" .}}`,
		"dashboard.html":   `{{define "content"}}{{template "content-dashboard" .}}{{end}}{{template "layout.html" .}}`,
		"user_detail.html": `{{define "content"}}{{template "content-user_detail" .}}{{end}}{{template "layout.html" .}}`,
		"error.html":       `{{define "content"}}{{template "content-error" .}}{{end}}{{template "layout.html" .}}`,
	}

	for name, content := range wrapperTemplates {
		template.Must(all.New(name).Parse(content))
	}

	r.SetHTMLTemplate(all)
	r.Static("/assets", "./assets")
}
