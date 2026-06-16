---
target: templates/index.html (landing page)
total_score: 25
p0_count: 0
p1_count: 2
p2_count: 3
timestamp: 2026-06-16T07-27-21Z
slug: templates-index-html
---
## Design Health Score

| # | Heuristic | Score | Key Issue |
|---|-----------|-------|-----------|
| 1 | Visibility of System Status | 3 | Scroll-based reveal animations work well, but no loading states for async quiz data |
| 2 | Match System / Real World | 3 | Indonesian language throughout is good; "Dark Triad" term may need inline definition for first-timers |
| 3 | User Control and Freedom | 3 | Nav is clear; mobile menu has close button; FAQ accordion allows collapse |
| 4 | Consistency and Standards | 3 | Design system is consistent; some backward-compat stubs in CSS create noise |
| 5 | Error Prevention | 2 | No form validation visible on landing; quiz page not reviewed |
| 6 | Recognition Rather Than Recall | 3 | Navigation is visible; icons have text labels |
| 7 | Flexibility and Efficiency | 1 | No keyboard shortcuts; no power user paths; single flow only |
| 8 | Aesthetic and Minimalist Design | 3 | Clean overall; some decorative icon usage violates "substance over decoration" |
| 9 | Error Recovery | 2 | No error states visible on landing; quiz error handling unknown |
| 10 | Help and Documentation | 2 | FAQ section exists but no contextual help; no tooltips |
| **Total** | | **25/40** | **Acceptable** |

## Anti-Patterns Verdict

**LLM assessment**: This does NOT immediately scream "AI made this" — and that's a genuine achievement. The design system is well-considered, the palette is intentional, and the restraint is real. However, there are several tells that push it toward the "refined but generic" zone rather than "distinctive brand":

1. **Inter as the sole typeface** — Inter is on the reflex-reject list for brand surfaces. While the single-family approach is deliberate per DESIGN.md, for a brand register it reads as safe rather than distinctive. The brand personality is "Dalam · Sunyi · Jujur" — Inter is none of those things; it's a neutral workhorse.

2. **Material Symbols icons** — The `material-symbols-outlined` icon set is used extensively (psychology, schedule, verified, lock, play_arrow, menu, close, home, info, check_circle, auto_awesome, bolt, assignment). For a brand that claims "substance over decoration," many of these icons are decorative rather than functional. The hero section's three metadata icons (schedule, verified, lock) are particularly tell-like — this is the "hero-metric template" variant.

3. **The hero section structure** — Centered headline + subtext + two buttons + three feature badges + sample card below. This is a well-executed but structurally generic hero layout. The "±5 Menit / Terpercaya / Data Aman" badges are the hero-metric template in miniature.

4. **Section heading pattern** — Every section has the same structure: centered heading + subtext paragraph. While not uppercase tracked eyebrows, the uniformity is a tell.

5. **Identical card grids in testimoni** — Three testimonials in identical card layouts with avatar + name + role + quote. This is the "identical card grid" anti-pattern.

**Deterministic scan**: The detector returned 0 findings — clean. No gradient text, no box-shadow abuse, no side-stripe borders, no glassmorphism outside the nav. The CSS is disciplined.

**Visual overlays**: Not available — the Go server requires DATABASE_URL and could not start. No browser visualization was possible.

## Overall Impression

ShadowSelf's landing page is a well-executed, disciplined implementation of a restrained design system. It successfully avoids the most egregious AI slop (no gradients, no glassmorphism, no illustrations, no uppercase eyebrows). The flat design, tonal layering, and generous whitespace create a calm, serious tone that aligns with the brand's "Dalam · Sunyi · Jujur" personality.

The single biggest opportunity: the page is **safe** when the brand brief calls for **distinctive**. "Sunyi" (quiet) doesn't mean invisible — and right now the page errs toward invisible. For a brand about confronting dark truths about oneself, the visual experience is almost too comfortable. There's no tension, no edge, no moment that makes the visitor pause. The teal accent is used sparingly (good), but nothing else creates emotional friction. A personality assessment about dark personality traits should feel slightly unsettling, not like a SaaS landing page.

## What's Working

1. **Disciplined design system adherence** — The CSS variables, spacing scale, component tokens, and flat-by-default philosophy are consistently applied. No shadows, no gradients, no decorative fluff. The design system document is excellent and the code follows it.

2. **Accessibility foundations** — Focus-visible outlines, reduced motion support, sr-only utility, semantic HTML (role attributes, aria labels), and contrast-aware color choices are all present. The `prefers-reduced-motion: reduce` media query is correctly implemented.

3. **FAQ accordion interaction** — The Alpine.js-powered accordion with x-collapse, proper aria-expanded/aria-controls, and the rotating "+" icon is well-implemented and accessible.

## Priority Issues

### [P1] The brand feels generic despite the design system

**What**: The landing page could belong to any introspective SaaS product. The hero layout (centered headline + two buttons + three feature badges + preview card) is structurally identical to thousands of other landing pages. The brand personality "Dalam · Sunyi · Jujur" doesn't translate into a visual experience that feels specific to ShadowSelf.

**Why it matters**: For a brand-register project where "design IS the product," a generic layout undermines trust. Users seeking answers about dark personality traits need to feel they've arrived somewhere authoritative and specific, not a template.

**Fix**: Introduce a distinctive visual anchor — a typographic treatment, a layout asymmetry, a dark/light section contrast that creates tension, or a subtle texture that evokes "depth" (the "Dalam" quality). The hero could be split into two columns (text left, something evocative right) rather than centered. The color strategy could shift from "restrained" to "committed" — use the deep teal as a section background rather than just an accent.

**Suggested command**: `$impeccable bolder`

### [P1] Inter on the reflex-reject list for brand surfaces

**What**: Inter is the sole typeface. Per the brand register's reflex-reject list, Inter is a training-data default that creates monoculture. For a brand about depth and darkness, Inter is a neutral, safe choice.

**Why it matters**: Typography is the primary carrier of brand voice in this system. A more distinctive typeface would immediately differentiate ShadowSelf from the thousands of other Inter-powered sites. The brand words "Dalam · Sunyi · Jujur" deserve a typeface with more character.

**Fix**: Replace Inter with a typeface that has more personality while maintaining readability. Options: a grotesque with more edge (like ABC Diatype, Untitled Sans, or even a variable font with optical size control), or pair Inter with a distinctive display face for headings only. JetBrains Mono for data is fine — keep that.

**Suggested command**: `$impeccable typeset`

### [P2] Material Symbols icons are decorative, not functional

**What**: The landing page uses Material Symbols extensively: psychology icon in logo, schedule/verified/lock in hero badges, auto_awesome/psychology/bolt for dimension cards, check_circle for feature list, assignment for the CTA section. Many of these are decorative rather than informative.

**Why it matters**: Design principle #2 is "Substansi di atas dekorasi" (Substance over decoration). Icons that don't carry meaning are decoration. The hero badges (schedule, verified, lock) are particularly questionable — they're the "hero-metric template" in miniature form.

**Fix**: Remove decorative icons. Keep only icons that serve a functional purpose (like the menu toggle, FAQ expand/collapse). Replace the hero badges with plain text. For the dimension cards, use typographic treatments (large initial letters, or the dimension name styled distinctively) instead of icons.

**Suggested command**: `$impeccable distill`

### [P2] No emotional tension or "darkness" in the visual design

**What**: The page is entirely light mode: white surface, light gray backgrounds, near-black text. For a brand called "ShadowSelf" that deals with dark personality traits, the visual experience is bright and airy. There's no shadow (literally — the design system bans it), no darkness, no weight.

**Why it matters**: The brand personality includes "Dalam" (deep) and the landing page should "evoke rasa waspada yang tenang" (evoke calm vigilance). A bright white page with light gray cards does not evoke vigilance or depth. It evokes a clean blog.

**Fix**: Introduce a dark section or dark hero variant. Use the deep teal or near-black as a background for at least one major section. Create visual contrast between "light" (the conscious self) and "dark" (the shadow self) through alternating dark/light section backgrounds. This would be a "Committed" color strategy move.

**Suggested command**: `$impeccable colorize`

### [P2] Testimonial cards are identical and feel templated

**What**: Three testimonials in identical card layouts: avatar circle + name + role + quote. The layout is the same for all three, staggered only by animation delay.

**Why it matters**: This is the "identical card grid" anti-pattern. Real testimonials should feel organic, not templated. The uniformity undermines authenticity — these read as placeholder content.

**Fix**: Vary the testimonial presentation. One could be a pull quote (large, typographic), another a short highlight with the score, another a more narrative format. Break the grid — let one testimonial be wider or use a different background treatment.

**Suggested command**: `$impeccable layout`

### [P3] CSS contains backward-compatible stubs and dead code

**What**: The CSS file contains ~40 lines of backward-compatible stubs (`.card-hover`, `.card-elevated`, `.shadow-glow`, `.bg-gradient-*`, `.bg-dots`, `.text-gradient`, `.hover-scale`, `.animate-spin`, `.pulse-dot`, etc.) that are either empty overrides or unused on the landing page.

**Why it matters**: Dead code creates maintenance burden and makes the actual design system harder to read. The `.shadow-glow { box-shadow: none; }` and `.bg-gradient-hero { background: var(--surface); }` stubs suggest these were once gradient/shadow elements that were overridden rather than removed.

**Fix**: Audit and remove unused CSS classes. Keep only what's actively used. Move any genuinely needed backward-compat stubs to a separate legacy file.

**Suggested command**: `$impeccable polish`

## Persona Red Flags

### Jordan (Confused First-Timer)

- **"Dark Triad" is not defined on first mention** — The hero paragraph says "asesmen psikologi berbasis Dark Triad" but doesn't explain what Dark Triad means until the FAQ section (scroll 3+ screens down). A first-timer may not know this term.
- **No visible "how it works" step-through** — The page explains WHAT the assessment measures but not HOW the process works (how many questions, what the experience feels like, what happens after). Jordan needs this reassurance before clicking "Mulai Tes."
- **The sample report card shows scores (75%, 60%, 35%)** — Without context, a first-timer might misinterpret these as their own scores or worry about what "high" means. The preview could cause anxiety rather than curiosity.

### Casey (Distracted Mobile User)

- **Primary CTA ("Mulai Tes") is in the top-right on mobile** — The mobile nav puts the CTA in the header, which is thumb-reachable. Good. But the hero CTA buttons are centered mid-page, requiring scroll and precise tap.
- **No state persistence mentioned** — If Casey starts the quiz, gets interrupted, and returns, will their progress be saved? No indication.
- **The page is content-heavy for mobile** — Hero + sample card + 3 dimensions + feature list + 3 testimonials + 4 FAQs + footer. That's a lot of scrolling for a user who may have limited attention.

### Riley (Deliberate Stress Tester)

- **What happens if JavaScript is disabled?** — Alpine.js handles the FAQ accordion and mobile menu. Without JS, the FAQ content is hidden (x-show without a noscript fallback) and the mobile menu button does nothing.
- **The `.fade-in-up` animation gates content visibility** — Elements start at `opacity: 0` and become visible only when the IntersectionObserver fires. If the observer fails or the user has a slow connection, content remains invisible. The JS has a fallback for missing IntersectionObserver, but not for observer failure.
- **Anchor links use smooth scroll via JS** — The app.js intercepts all `href^="#"` clicks. If JS fails, the browser's native anchor behavior should still work (since the href is valid), but the smooth scroll won't.

## Minor Observations

1. **The "Preview" badge on the sample report card** uses `rounded-full` (pill shape) — the design system says pill shapes are for progress bars only. Badges should use `rounded.sm` (8px) or `rounded.md` (12px).
2. **The footer has three metadata badges** (Berbasis Ilmiah, Data Terenkripsi, Hasil Instan) with teal icons — these repeat the hero badges' pattern and add clutter to what should be a minimal footer.
3. **The `pt-14` on `<main>`** (56px) is hardcoded rather than using a CSS variable. If the nav height changes, this breaks.
4. **The `max-w-[1120px]`** in the navbar uses arbitrary value syntax rather than the container's `max-width: 1120px`. Minor inconsistency.
5. **The `style="color: var(--inkMuted)"`** inline styles are used extensively instead of utility classes like `.text-muted`. This works but makes the HTML harder to maintain.
6. **The `x-cloak` directive** is present but the CSS rule `[x-cloak] { display: none !important; }` is at the bottom of the stylesheet — it should be loaded first to prevent FOUC.
7. **No `lang="id"` on the `<html>` tag** — wait, it IS there in layout.html. Good.
8. **The `impeccable-live` script tags** are still in layout.html (lines 22-24) — these are development artifacts that should be removed in production.

## Questions to Consider

1. **"What if the hero had a dark mode?"** — A dark hero section (near-black background, white text, teal accent) would immediately signal "ShadowSelf" and create the emotional tension the brand needs. The contrast between dark hero and light content sections would visually enact the "shadow self" concept.

2. **"Does this page need to be this long?"** — Hero + sample card + 3 dimensions + feature list + 3 testimonials + 4 FAQs + footer. For a brand about "Dalam" (depth), a shorter, more focused page with deeper individual sections might serve the brand better than a checklist of standard landing page sections.

3. **"What would make this feel like a destination rather than a tool?"** — Right now it reads as "here's a quiz, take it." A more editorial approach — a long-form essay about the shadow self, an interactive element, a thought-provoking question that scrolls into the assessment — would make the page feel like an experience rather than a funnel.

4. **"Is the 'no shadows' rule serving the brand name?"** — The design system bans box-shadow entirely. But the brand is called "ShadowSelf." A complete absence of shadow feels like a missed opportunity. Even a subtle shadow on interactive elements could reinforce the brand name without violating the flat aesthetic.
