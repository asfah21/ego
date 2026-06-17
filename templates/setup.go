package templates

import (
	"html/template"
	"math"
	"strings"
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

	// 1. Parse sections (navbar, footer, sidebar, topbar, produk_section, faq_section, testimoni_section, etc.)
	all := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/sections/*.html"))

	// 2. Parse all layouts
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/layout-*.html"))

	// 3. Parse index.html — uses {{define "content-index"}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/index.html"))

	// 4. Parse each page file individually — each uses {{define "content-*"}}
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
			if t.Name() != "" && strings.HasPrefix(t.Name(), "content-") {
				all = template.Must(all.AddParseTree(t.Name(), t.Tree))
			}
		}
	}

	// 5. Create wrapper templates named after each page file so Gin's c.HTML() can find them.
	//    Each wrapper:
	//      - defines {{define "content"}} that calls the specific content-* template
	//      - then calls the appropriate layout template which renders the full HTML structure
	//
	//    Layout mapping:
	//      layout-public.html    → navbar + footer (index, tentang, paywall, hasil, error)
	//      layout-quiz.html      → focused, no navbar/footer (quiz)
	//      layout-auth.html      → no navbar/footer, centered (login)
	//      layout-dashboard.html → sidebar + topbar (dashboard, user_detail)
	wrapperTemplates := map[string]string{
		// Public pages — navbar + footer
		"index.html":   `{{define "content"}}{{template "content-index" .}}{{end}}{{template "layout-public.html" .}}`,
		"tentang.html": `{{define "content"}}{{template "content-tentang" .}}{{end}}{{template "layout-public.html" .}}`,
		"paywall.html": `{{define "content"}}{{template "content-paywall" .}}{{end}}{{template "layout-public.html" .}}`,
		"hasil.html":   `{{define "content"}}{{template "content-hasil" .}}{{end}}{{template "layout-public.html" .}}`,
		"error.html":   `{{define "content"}}{{template "content-error" .}}{{end}}{{template "layout-public.html" .}}`,

		// Quiz page — focused, no public navbar/footer
		"quiz.html": `{{define "content"}}{{template "content-quiz" .}}{{end}}{{template "layout-quiz.html" .}}`,

		// Auth page — no navbar/footer, centered
		"login.html": `{{define "content"}}{{template "content-login" .}}{{end}}{{template "layout-auth.html" .}}`,

		// Dashboard pages — sidebar + topbar
		"dashboard.html":   `{{define "content"}}{{template "content-dashboard" .}}{{end}}{{template "layout-dashboard.html" .}}`,
		"user_detail.html": `{{define "content"}}{{template "content-user_detail" .}}{{end}}{{template "layout-dashboard.html" .}}`,
	}

	for name, content := range wrapperTemplates {
		template.Must(all.New(name).Parse(content))
	}

	r.SetHTMLTemplate(all)
	r.Static("/assets", "./assets")
}
