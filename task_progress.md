# Task Progress: Fix Template Rendering (DOCTYPE, Layout, Assets)

## Root Cause Analysis

**The layout files (`layout-public.html`, `layout-auth.html`, `layout-dashboard.html`, `layout-quiz.html`) are NEVER used.**

In `templates/setup.go`, the wrapper templates directly render sections (navbar, footer, sidebar, topbar) and content blocks inline, completely bypassing the layout files. The layout files contain the critical `<!DOCTYPE html>`, `<html>`, `<head>`, `<body>` structure, CSS/JS/font links, etc.

The wrapper templates in `setup.go` only render:
- `{{template "navbar" .}}<main>{{template "content-*" .}}</main>{{template "footer" .}}`

This means the rendered HTML is just a fragment like:
```html
<header class="navbar">...<main>...content...</main><footer>...</footer>
```

**No DOCTYPE, no `<html>`, no `<head>`, no `<body>`, no CSS/JS/font links.**

## Fix Strategy

Modify the wrapper templates in `templates/setup.go` to include the full HTML structure from the layout files. Each wrapper template must include:
1. `<!DOCTYPE html>`
2. `<html>`
3. `<head>` with all CSS, JS, fonts, meta tags
4. `<body>` with the appropriate class
5. The page-specific sections (navbar, footer, sidebar, topbar)
6. The content block

I'll create a shared "head" section that contains all the common `<head>` content, then reference it from each wrapper template.
