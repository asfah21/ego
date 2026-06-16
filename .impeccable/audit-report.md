# ShadowSelf — Design Audit Report

**Date:** 2026-06-16  
**Audit type:** Carbonization — gradient eradication & typography refinement  
**Scope:** All templates, CSS design system  
**Detector result:** Clean (0 findings)

---

## 1. Changes Made

### Palette Carbonization

| File | Change | Status |
|------|--------|--------|
| `assets/css/style.css` | Removed `--warmBg`, `--warmBord` variables; replaced with `--bg`, `--bord` | ✅ |
| `assets/css/style.css` | Removed `--warm-50` through `--warm-600` (replaced with semantic aliases) | ✅ |
| `assets/css/style.css` | Replaced all `--warmBg` references with `var(--surface)` or `var(--bg)` | ✅ |
| `assets/css/style.css` | Replaced all `--warmBord` references with `var(--bordLight)` | ✅ |
| `assets/css/style.css` | Added `--accentLight` and `--accentBorder` for accent backgrounds | ✅ |

### Gradient Eradication

| File | Change | Status |
|------|--------|--------|
| `templates/_pages/quiz.html` | Replaced `linear-gradient(90deg,#10b981,#059669)` with `var(--accent)` | ✅ |
| `templates/_pages/quiz.html` | Replaced loading bar gradient with `var(--accent)` | ✅ |
| `templates/_pages/hasil.html` | Replaced Narsisme gradient with `var(--narcissus)` | ✅ |
| `templates/_pages/hasil.html` | Replaced Machiavellian gradient with `var(--machiavellian)` | ✅ |
| `templates/_pages/hasil.html` | Replaced Psikopati gradient with `var(--psychopath)` | ✅ |
| `templates/_pages/hasil.html` | Replaced `bg-gradient-to-r` with solid `var(--accentLight)` background | ✅ |
| `templates/_pages/hasil.html` | Replaced `text-gradient` with `text-accent` | ✅ |
| `templates/_pages/tentang.html` | Replaced `text-gradient` with `text-accent` | ✅ |
| `templates/_pages/tentang.html` | Removed `shadow-glow` class | ✅ |

### Typography — DM Serif Display

| File | Change | Status |
|------|--------|--------|
| `templates/layout.html` | Added DM Serif Display font import | ✅ |
| `assets/css/style.css` | Added `h1, h2` rule: `font-family: 'DM Serif Display'` | ✅ |
| `assets/css/style.css` | Removed redundant `font-family: Inter` from h1/h2 base rule | ✅ |

### Navbar CTA

| File | Change | Status |
|------|--------|--------|
| `assets/css/style.css` | Changed `.navbar-cta` background from `var(--ink)` to `var(--accent)` | ✅ |

---

## 2. Design System Compliance

| Rule | Status | Notes |
|------|--------|-------|
| Flat by default (no box-shadow) | ✅ | No shadows introduced |
| One accent (Abyssal Teal) | ✅ | Teal used in CTA, icons, accent text |
| Data-only colors | ✅ | Narcissus/Machiavellian/Psychopath only in progress bars |
| No gradient text | ✅ | All `text-gradient` → `text-accent` (solid color) |
| No gradient backgrounds | ✅ | All `linear-gradient` → solid CSS variables |
| No decorative gradients | ✅ | `bg-gradient-to-r` → solid `var(--accentLight)` |
| No glassmorphism outside nav | ✅ | Only navbar uses backdrop-filter |
| No illustrations | ✅ | None |
| No uppercase eyebrows | ✅ | None |
| No numbered section markers | ✅ | None |
| Border-radius ≤ 12px on cards | ✅ | All cards use 12px |
| Pill shapes for progress bars only | ✅ | Progress bars use `rounded-full` |
| DM Serif Display for h1/h2 | ✅ | Editorial gravity for headings |

---

## 3. Anti-Pattern Scan

### Detector Results: **Clean (0 findings)**

The deterministic scanner found no instances of:
- Gradient text
- Box-shadow abuse
- Side-stripe borders
- Glassmorphism outside nav
- Decorative patterns
- Uppercase tracked eyebrows
- `linear-gradient` in templates
- `bg-gradient-to-r` in templates
- `shadow-glow` in templates

---

## 4. Summary

| Area | Before | After | Delta |
|------|--------|-------|-------|
| Warm-tinted palette | `--warmBg`, `--warmBord` | Carbon-neutral `--bg`, `--bord` | ✅ |
| Gradient progress bars | 5 instances | 0 — all solid colors | ✅ |
| Gradient text | 2 instances | 0 — all `text-accent` | ✅ |
| Gradient backgrounds | 1 instance | 0 — solid `var(--accentLight)` | ✅ |
| `shadow-glow` usage | 1 instance | 0 — removed | ✅ |
| Heading typeface | Inter only | DM Serif Display + Inter | ✅ |
| Navbar CTA color | Black (`var(--ink)`) | Teal (`var(--accent)`) | ✅ |
| Detector findings | 0 | 0 | ✅ |
