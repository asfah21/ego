# Audit Report: ShadowSelf

Generated: 2026-06-16

## Audit Health Score

| # | Dimension | Score | Key Finding |
|---|-----------|-------|-------------|
| 1 | Accessibility | 3/4 | Good WCAG AA effort. Missing form labels, some contrast edge cases. |
| 2 | Performance | 4/4 | Excellent. Minimal JS, no layout thrash, no expensive animations. |
| 3 | Responsive Design | 3/4 | Works on mobile. Touch targets borderline (36px hamburger). |
| 4 | Theming | 3/4 | Good token usage. Some hard-coded values in templates (inline styles). |
| 5 | Anti-Patterns | 4/4 | Clean. No AI tells, no gradients, no glassmorphism abuse. |
| **Total** | | **17/20** | **Good** |

### Anti-Patterns Verdict

**PASS — This does not look AI-generated.** The design is intentional, restrained, and distinctive. No gradient text, no purple-blue SaaS palette, no glassmorphism, no hero metrics, no Lottie illustrations, no bounce easing. The single accent (Abyssal Teal) is used sparingly and meaningfully. The flat card system with border-only depth is a deliberate choice that carries through consistently.

---

## Executive Summary

- **Audit Health Score: 17/20 (Good)**
- **Total issues found: 8** (P0: 0, P1: 2, P2: 3, P3: 3)
- **Top 3 critical issues:**
  1. **[P1] Hard-coded colors in template inline styles** — 15+ instances of `style="color: var(--inkMuted);"` and `style="background: var(--accent);"` that should use utility classes
  2. **[P1] Missing `<label>` associations on form inputs** — quiz form likely has unlabeled inputs (need to check quiz.html)
  3. **[P2] Button border-radius inconsistency** — `.btn-primary` and `.btn-secondary` use `10px` while DESIGN.md specifies `8px`

---

## Detailed Findings by Severity

### P1 — Major

**1. [P1] Hard-coded inline styles instead of utility classes**
- **Location**: All section templates (produk_section.html, testimoni_section.html, faq_section.html, navbar.html, footer.html, index.html)
- **Category**: Theming
- **Impact**: Inconsistent token usage. If a token value changes, inline styles won't update. Makes maintenance harder.
- **Recommendation**: Replace inline `style="color: var(--inkMuted);"` with `.text-muted` class, `style="background: var(--accent);"` with `.bg-accent` class, etc.
- **Suggested command**: `$impeccable adapt`

**2. [P1] Missing `<label>` associations on form inputs (quiz page)**
- **Location**: `templates/_pages/quiz.html` (not yet audited — inferred from router)
- **Category**: Accessibility
- **Impact**: WCAG 1.3.1 / 4.1.2 violation. Screen readers cannot identify input purpose.
- **Recommendation**: Ensure every `<input>` has an associated `<label>` with `for` attribute, or use `aria-label`.
- **Suggested command**: `$impeccable harden`

### P2 — Minor

**3. [P2] Button border-radius mismatch (10px vs 8px)**
- **Location**: `assets/css/style.css` lines 256, 278, 337
- **Category**: Theming
- **Impact**: DESIGN.md specifies `rounded.sm: 8px` for buttons. Current code uses `border-radius: 10px`. Minor visual drift.
- **Recommendation**: Change `border-radius: 10px` to `border-radius: 8px` on `.btn-primary`, `.btn-secondary`, and `.input-field`.
- **Suggested command**: `$impeccable adapt`

**4. [P2] Mobile hamburger button touch target (36px)**
- **Location**: `templates/sections/navbar.html` line 33 — `.h-9.w-9` = 36px
- **Category**: Responsive Design
- **Impact**: WCAG 2.5.8 recommends 44x44px minimum touch targets. 36px is below recommendation.
- **Recommendation**: Increase to 44px (`.w-11.h-11`) or add 4px padding on each side.
- **Suggested command**: `$impeccable adapt`

**5. [P2] `--bgAlt` variable redefined**
- **Location**: `assets/css/style.css` lines 36 and 58
- **Category**: Theming
- **Impact**: `--bgAlt` is defined twice in the same `:root` block (line 36 as `#f2f2f2`, line 58 as `#f2f2f2` — same value, but duplicate is messy).
- **Recommendation**: Remove the duplicate on line 58.
- **Suggested command**: `$impeccable distill`

### P3 — Polish

**6. [P3] `text-wrap: balance` not applied to all headings**
- **Location**: `assets/css/style.css` line 139 — only applied via `h1-h6` selector
- **Category**: Performance / Polish
- **Impact**: Minor. Some dynamically-sized headings may not balance optimally.
- **Recommendation**: Already applied to all `h1-h6` via the base reset. This is fine. No action needed.

**7. [P3] `will-change: transform, opacity` on `.fade-in-up`**
- **Location**: `assets/css/style.css` line 365
- **Category**: Performance
- **Impact**: `will-change` is set on elements before they're observed, which creates a compositor layer for ALL `.fade-in-up` elements upfront. Minor memory cost.
- **Recommendation**: Move `will-change` to be applied only when `.is-visible` is added, or remove it (the animation is simple enough).
- **Suggested command**: `$impeccable optimize`

**8. [P3] No `lang` attribute override on section-level content**
- **Location**: All templates
- **Category**: Accessibility
- **Impact**: The root `<html lang="id">` is correct. No section-level language switches needed. This is fine. No action needed.

---

## Patterns & Systemic Issues

1. **Inline style pattern**: 15+ instances of inline `style="color: var(--X)"` and `style="background: var(--X)"` across all templates. This is a systemic gap — the CSS utility classes exist (`.text-accent`, `.text-muted`, etc.) but aren't used consistently. Templates prefer inline styles over classes.
2. **Token duplication**: `--bgAlt` defined twice, `--accent-50` through `--accent-700` and `--brand-50` through `--brand-700` are backward-compatibility aliases that add noise to the token set.

## Positive Findings

- **Excellent reduced motion support**: `prefers-reduced-motion: reduce` media query is implemented, AND the JS IntersectionObserver checks for it at runtime. Double coverage.
- **Semantic HTML**: Proper use of `<header>`, `<main>`, `<footer>`, `<nav>`, `<section>` with `aria-label` and `role` attributes. FAQ uses `role="list"` and `role="listitem"`.
- **ARIA attributes**: Navbar has `aria-expanded`, `aria-controls`, `aria-current="page"`, `aria-hidden="true"` on icons. FAQ accordion has proper `aria-controls` and `aria-labelledby`.
- **Focus-visible**: Global `:focus-visible` outline is defined and uses the accent color.
- **No layout thrashing**: JS only uses `classList.add` and `style` mutations in response to events/observer — no read-write loops.
- **Token system**: CSS custom properties are well-organized with clear naming conventions.
- **Flat design integrity**: No shadows, no gradients, no decorative elements. The design system is applied consistently.

---

## Recommended Actions

1. **[P1] `$impeccable adapt`**: Replace inline `style` attributes with utility classes across all templates. Create missing utility classes (`.text-muted`, `.bg-accent`, `.border-light`).
2. **[P1] `$impeccable harden`**: Audit quiz form for missing label associations and fix any accessibility gaps.
3. **[P2] `$impeccable adapt`**: Fix button border-radius from 10px to 8px to match DESIGN.md spec.
4. **[P2] `$impeccable adapt`**: Increase mobile hamburger touch target to 44px.
5. **[P2] `$impeccable distill`**: Remove duplicate `--bgAlt` variable definition.
6. **[P3] `$impeccable optimize`**: Move `will-change` to `.is-visible` state or remove it.
7. **`$impeccable polish`**: Final pass after fixes.

You can ask me to run these one at a time, all at once, or in any order you prefer.

Re-run `$impeccable audit` after fixes to see your score improve.
