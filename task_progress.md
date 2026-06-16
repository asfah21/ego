# ShadowSelf — Impeccable Polish

## Completed

- [x] **CSS Architecture** — Consolidated and reorganized `style.css` with clear sections (variables, reset, typography, containers, sections, cards, buttons, progress bars, inputs, glass, animations, utilities, responsive, backward-compat stubs)
- [x] **Animation System** — Replaced inline `opacity`/`transform` with `.fade-in-up` + `.is-visible` pattern; JS updated to use `.is-visible` class; `prefers-reduced-motion` respected
- [x] **Inline Style Migration** — Replaced `style="color: var(--ink)"` with `.text-ink`, `style="color: var(--accent)"` with `.text-accent`, etc. across all templates
- [x] **Accessibility (ARIA)** — Added `role="banner"`, `role="contentinfo"`, `role="list"`, `role="listitem"`, `aria-expanded`, `aria-controls`, `aria-label`, `aria-labelledby`, `aria-hidden="true"` on icons, `aria-current="page"` on nav
- [x] **Semantic HTML** — Added `aria-labelledby` to FAQ section, `id` attributes to headings, `role="region"` to FAQ panels
- [x] **Utility Classes** — Added missing utilities: `.text-ink`, `.leading-normal`, `.tracking-tight`, `.z-50`, `.fixed`, `.top-0`, `.inline-flex`, `.pt-14`, `.px-2`, `.px-2.5`, `.py-0.5`
- [x] **Code Cleanup** — Removed inline `[x-cloak]` style from layout.html (now in CSS), removed unused `m-0` classes from elements that already have `p { margin: 0 }` reset
- [x] **Build Verification** — Go project compiles successfully
