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

	// 1. Parse sections (navbar, footer, produk_section, etc.)
	all := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/sections/*.html"))

	// 2. Parse layout — defines outer HTML structure, calls {{template "content" .}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/layout.html"))

	// 3. Parse index.html — uses {{define "content"}}
	all = template.Must(template.Must(all.Clone()).ParseGlob("templates/index.html"))

	// 4. Parse each page file individually — each uses a unique define name
	//    e.g. {{define "content-quiz"}}, {{define "content-login"}}, etc.
	//    IMPORTANT: Only add templates whose names start with "content-" to avoid
	//    overwriting existing templates (like "quiz.html", "paywall.html", etc.)
	//    that were already parsed from sections, layout, and index.html.
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
	//    Each wrapper renders the full page.
	//    IMPORTANT: We must NOT use {{define "content"}} in wrapper templates because that would
	//    overwrite the global "content" template (defined in index.html) for ALL pages.
	//    Instead, each wrapper is a self-contained page that includes navbar, the specific
	//    content template, and footer directly.
	//    For index.html, we use layout.html which calls {{template "content" .}} (defined in index.html).
	wrapperTemplates := map[string]string{
		"index.html":       `{{template "layout.html" .}}`,
		"quiz.html":        `{{template "navbar" .}}<main>{{template "content-quiz" .}}</main>{{template "footer" .}}`,
		"paywall.html":     `{{template "navbar" .}}<main>{{template "content-paywall" .}}</main>{{template "footer" .}}`,
		"hasil.html":       `{{template "navbar" .}}<main>{{template "content-hasil" .}}</main>{{template "footer" .}}`,
		"tentang.html":     `{{template "navbar" .}}<main>{{template "content-tentang" .}}</main>{{template "footer" .}}`,
		"login.html":       `{{template "navbar" .}}<main>{{template "content-login" .}}</main>{{template "footer" .}}`,
		"dashboard.html":   `{{template "navbar" .}}<main>{{template "content-dashboard" .}}</main>{{template "footer" .}}`,
		"user_detail.html": `{{template "navbar" .}}<main>{{template "content-user_detail" .}}</main>{{template "footer" .}}`,
		"error.html":       `{{template "navbar" .}}<main>{{template "content-error" .}}</main>{{template "footer" .}}`,
	}

	for name, content := range wrapperTemplates {
		template.Must(all.New(name).Parse(content))
	}

	r.SetHTMLTemplate(all)
	r.Static("/assets", "./assets")
}
