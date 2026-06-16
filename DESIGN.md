---
name: ShadowSelf
description: Dark Triad personality assessment — quiet, honest, introspective
colors:
  primary: "#0d7377"
  primary-hover: "#095b5e"
  primary-light: "#e8f3f4"
  primary-border: "#c5e0e2"
  ink: "#0c0c0c"
  ink-muted: "#555555"
  ink-subtle: "#999999"
  bord: "#e5e5e5"
  bord-light: "#f0f0f0"
  surface: "#ffffff"
  bg: "#f8f8f8"
  bg-alt: "#f2f2f2"
  narcissus: "#7c3aed"
  machiavellian: "#d97706"
  psychopath: "#059669"
  success: "#059669"
  warning: "#d97706"
  error: "#dc2626"
typography:
  display:
    fontFamily: "Inter, system-ui, -apple-system, sans-serif"
    fontSize: "clamp(2.25rem, 5vw, 4rem)"
    fontWeight: 800
    lineHeight: 1.1
    letterSpacing: "-0.03em"
  headline:
    fontFamily: "Inter, system-ui, -apple-system, sans-serif"
    fontSize: "clamp(1.75rem, 3vw, 2.5rem)"
    fontWeight: 700
    lineHeight: 1.1
    letterSpacing: "-0.03em"
  title:
    fontFamily: "Inter, system-ui, -apple-system, sans-serif"
    fontSize: "clamp(1.5rem, 2.5vw, 2rem)"
    fontWeight: 700
    lineHeight: 1.1
    letterSpacing: "-0.03em"
  body:
    fontFamily: "Inter, system-ui, -apple-system, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.65
    letterSpacing: "0em"
  label:
    fontFamily: "Inter, system-ui, -apple-system, sans-serif"
    fontSize: "0.875rem"
    fontWeight: 500
    lineHeight: 1.5
    letterSpacing: "0em"
  mono:
    fontFamily: "JetBrains Mono, monospace"
    fontSize: "0.875rem"
    fontWeight: 600
    lineHeight: 1.5
    letterSpacing: "0em"
rounded:
  sm: "8px"
  md: "12px"
  full: "9999px"
spacing:
  "2xs": "4px"
  xs: "8px"
  sm: "12px"
  md: "16px"
  lg: "24px"
  xl: "32px"
  "2xl": "48px"
  "3xl": "64px"
  "4xl": "80px"
  "5xl": "96px"
  "6xl": "128px"
components:
  button-primary:
    backgroundColor: "{colors.ink}"
    textColor: "#ffffff"
    rounded: "{rounded.sm}"
    padding: "12px 28px"
    typography: "{typography.label}"
  button-primary-hover:
    backgroundColor: "#2a2a2a"
    textColor: "#ffffff"
    rounded: "{rounded.sm}"
    padding: "12px 28px"
  button-secondary:
    backgroundColor: "transparent"
    textColor: "{colors.ink-muted}"
    rounded: "{rounded.sm}"
    padding: "12px 28px"
    border: "1px solid {colors.bord}"
  button-secondary-hover:
    backgroundColor: "transparent"
    textColor: "{colors.ink}"
    rounded: "{rounded.sm}"
    padding: "12px 28px"
    border: "1px solid {colors.ink}"
  button-ghost:
    backgroundColor: "transparent"
    textColor: "{colors.ink-muted}"
    rounded: "{rounded.sm}"
    padding: "8px 16px"
  button-ghost-hover:
    backgroundColor: "{colors.bg}"
    textColor: "{colors.ink}"
    rounded: "{rounded.sm}"
    padding: "8px 16px"
  card:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.md}"
    border: "1px solid {colors.bord-light}"
  card-flat:
    backgroundColor: "{colors.bg}"
    rounded: "{rounded.md}"
    border: "none"
  input-field:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.ink}"
    rounded: "{rounded.sm}"
    padding: "12px 16px"
    border: "1px solid {colors.bord}"
  input-field-focus:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.ink}"
    rounded: "{rounded.sm}"
    padding: "12px 16px"
    border: "1px solid {colors.primary}"
  progress-bar:
    backgroundColor: "{colors.bord-light}"
    rounded: "{rounded.full}"
    height: "4px"
  progress-fill:
    rounded: "{rounded.full}"
    height: "100%"
---

# Design System: ShadowSelf

## 1. Overview

**Creative North Star: "The Still Mirror"**

ShadowSelf's visual system is a mirror that doesn't flatter and doesn't distort — it simply reflects what's there. The interface exists to create space for a difficult conversation with oneself, not to distract or entertain. Every pixel is in service of that stillness.

The system is **refined and restrained**: buttons don't shout, cards don't float, colors don't compete. The deep teal accent ("Abyssal Teal") appears sparingly — not as decoration, but as a quiet signal: *this is where to look, this is what matters.* The palette is intentionally narrow: near-black ink, cool gray neutrals, and one accent that carries the emotional weight of the brand.

This system explicitly rejects: gamified personality tests (16Personalities), corporate wellness warmth (BetterHelp), generic SaaS gradients, and decorative aesthetics that substitute substance for style.

**Key Characteristics:**
- Flat by default, no shadows — depth is conveyed through tonal layering and borders, not elevation
- Narrow, intentional palette — one accent, cool neutrals, no decorative colors
- Typography-forward — Inter at multiple weights carries the hierarchy, not color or decoration
- Generous whitespace — rhythm over density, room to breathe over information packing
- Still, not empty — every element has purpose; nothing exists "just because"

## 2. Colors

The palette is intentionally narrow. One accent carries the emotional weight; neutrals provide structure. Color is never decorative — it signals, it separates, it guides.

### Primary
- **Abyssal Teal** (#0d7377): The single accent. Used for interactive elements (focus rings, active states), data highlights, and the brand mark. Appears on ≤10% of any given screen. Its rarity is the point.
- **Abyssal Teal Hover** (#095b5e): Darkened for hover states on interactive elements.
- **Abyssal Teal Light** (#e8f3f4): Tinted for subtle backgrounds, badge fills, and icon containers.
- **Abyssal Teal Border** (#c5e0e2): The lightest reach of the accent, used for bordered containers that need a hint of brand.

### Neutral
- **Ink** (#0c0c0c): Primary text. Near-black, not pure black — softer on the eyes.
- **Ink Muted** (#555555): Secondary text, body copy. Meets WCAG AA 4.5:1 against white.
- **Ink Subtle** (#999999): Placeholder text, metadata, captions. Meets WCAG AA 4.5:1 against white.
- **Bord** (#e5e5e5): Standard borders, dividers.
- **Bord Light** (#f0f0f0): Subtle borders, card outlines, hairline dividers.
- **Surface** (#ffffff): Card backgrounds, page background.
- **Bg** (#f8f8f8): Section alternation, hover states, flat card backgrounds.
- **Bg Alt** (#f2f2f2): Deeper background alternation, subtle container fills.

### Data (Dark Triad)
- **Narcissus** (#7c3aed): Purple. The narcissism dimension. Used only in score displays and progress bars.
- **Machiavellian** (#d97706): Amber. The machiavellian dimension. Used only in score displays and progress bars.
- **Psychopath** (#059669): Emerald. The psychopathy dimension. Used only in score displays and progress bars.

### Semantic
- **Success** (#059669): Confirmation states, paid badges.
- **Warning** (#d97706): Paywall badges, cautionary indicators.
- **Error** (#dc2626): Error states, destructive actions.

### Named Rules

**The One Voice Rule.** Abyssal Teal is the only accent. It appears on ≤10% of any given screen. If a second accent color appears outside the Dark Triad data colors, the design has drifted.

**The Data-Only Rule.** Narcissus, Machiavellian, and Psychopath colors are reserved exclusively for score displays and progress bars. They never appear in navigation, buttons, or decorative elements.

## 3. Typography

**Display Font:** Inter (with system-ui, -apple-system, sans-serif fallback)
**Body Font:** Inter (same stack)
**Label/Mono Font:** JetBrains Mono (for scores, data points, code)

**Character:** A single-family system. Inter at 800 weight provides the gravity for display headings; at 400 it recedes into quiet readability for body text. The pairing is not a pairing — it's one voice at different registers. No serif, no contrast for contrast's sake. The weight shifts do the work.

### Hierarchy
- **Display** (800, clamp(2.25rem, 5vw, 4rem), 1.1): Hero headlines only. Max 6rem. Letter-spacing: -0.03em floor. `text-wrap: balance`.
- **Headline** (700, clamp(1.75rem, 3vw, 2.5rem), 1.1): Section headings. `text-wrap: balance`.
- **Title** (700, clamp(1.5rem, 2.5vw, 2rem), 1.1): Card titles, subheadings. `text-wrap: balance`.
- **Body** (400, 1rem, 1.65): Paragraphs, descriptions. Max line length 65–75ch. `text-wrap: pretty`.
- **Label** (500, 0.875rem, 1.5): Buttons, nav links, form labels, metadata.
- **Mono** (600, 0.875rem, 1.5): Score percentages, data values, progress labels.

### Named Rules

**The Weight-Only Hierarchy Rule.** Hierarchy is expressed through weight and size, not through color, tracking, or case changes. No uppercase eyebrows, no wide tracking on labels, no colored headings.

## 4. Elevation

The system is flat by default. No shadows, no drop shadows, no box-shadow on any surface. Depth is conveyed exclusively through:

- **Tonal layering**: Surface → Bg → Bg Alt creates a natural stacking order without elevation.
- **Borders**: 1px solid borders at `bordLight` or `bord` define container boundaries.
- **Hover state**: Border darkens from `bordLight` to `bord` on interactive cards. No lift, no shadow.

This is intentional. Shadows imply float; ShadowSelf is grounded. The interface sits flat on the page, like a printed psychological assessment — not a floating dashboard.

### Named Rules

**The Flat-By-Default Rule.** No surface casts a shadow. If an element needs to feel distinct, use tonal layering (surface → bg) or a border shift. `box-shadow` is prohibited.

## 5. Components

### Buttons

- **Shape:** Gently rounded (8px / `rounded.sm`). Not pill-shaped.
- **Primary:** Near-black (`ink`) background, white text, 12px 28px padding. Hover: darkens to #2a2a2a. No shadow, no glow.
- **Secondary:** Transparent background, `inkMuted` text, 1px `bord` border. Hover: border shifts to `ink`, text shifts to `ink`.
- **Ghost:** Transparent background, `inkMuted` text, 8px 16px padding. Hover: `bg` background, `ink` text.
- **States:** All buttons transition background/border-color at 0.2s ease. Focus-visible: 2px `accent` outline with 2px offset.

### Cards / Containers

- **Corner Style:** Moderately rounded (12px / `rounded.md`). Not pill-like.
- **Background:** `surface` (white) for elevated cards; `bg` (#f8f8f8) for flat cards.
- **Shadow Strategy:** None. Cards are defined by a 1px `bordLight` border.
- **Hover:** Border shifts to `bord`. No lift, no shadow.
- **Internal Padding:** 24px–32px (`space-lg` to `space-xl`), depending on content density.

### Inputs / Fields

- **Style:** 1px `bord` stroke, `surface` background, 8px radius.
- **Focus:** Border shifts to `accent` (#0d7377). No glow, no ring beyond the focus-visible outline.
- **Placeholder:** `inkSubtle` (#999999) — meets WCAG AA 4.5:1.
- **Error / Disabled:** Error: `error` border (#dc2626). Disabled: reduced opacity.

### Navigation

- **Style:** Fixed top bar, `glass` effect (rgba white 0.85 + 16px blur), 1px `bordLight` bottom border.
- **Links:** Ghost button style. Active page: no special indicator (no underline, no colored dot). The URL is the affordance.
- **Mobile:** Slide-down menu with fade + translate transition. Same visual language, stacked vertically.

### Progress Bars

- **Style:** 4px tall, `bordLight` track, `rounded.full` caps.
- **Fill:** Solid color matching the dimension (Narcissus purple, Machiavellian amber, Psychopath emerald). No gradients.
- **Transition:** Width animates at 0.8s ease.

## 6. Do's and Don'ts

### Do:
- **Do** use Abyssal Teal sparingly — ≤10% of any screen. Its rarity is its power.
- **Do** use tonal layering (surface → bg → bgAlt) to create depth instead of shadows.
- **Do** keep body text at `inkMuted` (#555555) or darker — meets WCAG AA 4.5:1.
- **Do** use `text-wrap: balance` on headings and `text-wrap: pretty` on body text.
- **Do** cap body line length at 65–75ch.
- **Do** use the Dark Triad colors (purple, amber, emerald) exclusively for score data.
- **Do** use Inter at different weights for hierarchy — weight shifts do the work.
- **Do** keep cards flat with a 1px `bordLight` border — no shadows, no lift.

### Don't:
- **Don't** gamify the interface — no badges, no achievements, no progress celebrations. This is not 16Personalities.
- **Don't** use corporate wellness warmth — no fake-warm tones, no stock photos of smiling people, no "you're great!" messaging. This is not BetterHelp.
- **Don't** use gradient backgrounds, gradient text, or purple-blue SaaS gradients anywhere.
- **Don't** use glassmorphism, blur overlays, or decorative backdrop filters outside the nav bar.
- **Don't** use illustrations — no Lottie, no SVG sketches, no hand-drawn elements. If you can't render it with type and layout, don't render it.
- **Don't** use uppercase tracked eyebrows ("ABOUT", "PROCESS") above section headings.
- **Don't** use numbered section markers (01, 02, 03) as default scaffolding.
- **Don't** use `box-shadow` on any surface — cards, buttons, or containers.
- **Don't** use border-radius larger than 12px on cards or containers. Pill shapes are for progress bars only.
- **Don't** pair border + shadow on the same element. Pick one.
- **Don't** use decorative patterns, repeating backgrounds, or diagonal stripes.
- **Don't** use the Dark Triad data colors outside of score displays.
