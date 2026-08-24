# Design System — hello-word-14

> Source of truth: the approved `index.html` (preview: approved design).
> Every value below is extracted from it. Changing a value here without
> changing the approved design is a defect.

Last updated: 2025-02-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#FFFFFF` | Page background |
| `--color-text` | `#000000` | Body text |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |

### 1.2 Spacing

Base unit: `4px`. Every margin, padding, and gap in the product uses one of these.

| Token | Value |
|---|---|
| `--space-1` | `4px` |
| `--space-2` | `8px` |
| `--space-3` | `12px` |
| `--space-4` | `16px` |
| `--space-6` | `24px` |
| `--space-8` | `32px` |
| `--space-12` | `48px` |

### 1.3 Typography

Font families (include the fallback stack and how the font is loaded):

- Body: `Arial, Helvetica, sans-serif` loaded by system fallback
- Headings: `Arial, Helvetica, sans-serif` loaded by system fallback
- Mono: `ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace` loaded by system fallback

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-xs` | `12px` | `1.4` | `400` | Caption, helper text |
| `--text-sm` | `14px` | `1.4` | `400` | Secondary body |
| `--text-base` | `16px` | `1.5` | `400` | Body |
| `--text-lg` | `18px` | `1.5` | `400` | Lead paragraph |
| `--text-xl` | `20px` | `1.2` | `400` | h3 |
| `--text-2xl` | `32px` | `1` | `400` | h2 |
| `--text-3xl` | `clamp(32px, 8vw, 80px)` | `1` | `400` | h1 / single-page hero text |

Heading levels are used in order and never skipped for visual sizing.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | `0px` | Input, badge |
| `--radius-md` | `0px` | Button, card |
| `--radius-lg` | `0px` | Modal |
| `--radius-full` | `9999px` | Avatar, pill |
| `--border-width` | `0px` | Default border |
| `--shadow-sm` | `none` | Resting card |
| `--shadow-md` | `none` | Dropdown, popover |
| `--shadow-lg` | `none` | Modal |
| `--duration-fast` | `0ms` | Hover, focus |
| `--duration-base` | `0ms` | Panel open/close |
| `--easing` | `linear` | All transitions |

Motion respects `prefers-reduced-motion: reduce`: state changes remain, movement is removed.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | `640px` | `100%` | `1` | `0px` |
| `md` | `768px` | `100%` | `1` | `0px` |
| `lg` | `1024px` | `100%` | `1` | `0px` |
| `xl` | `1280px` | `100%` | `1` | `0px` |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | `0` |
| Dropdown | `0` |
| Modal backdrop | `0` |
| Modal | `0` |
| Toast | `0` |

## 2. Components

One subsection per reusable component. Every component lists **all** states.

### 2.1 Centered Message

**Purpose** — Static full-page message, used only for single-line landing text. Not for interactive UI.

**Anatomy** — `[message text]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-text`, `--text-3xl` | Single centered page copy |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | Auto | `0` | `--text-3xl` |

**States** — every row must be filled in.

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered black text on white background | `--color-bg`, `--color-text`, `--text-3xl` |
| Hover | No change | None |
| Focus (keyboard) | No interactive target exists | None |
| Active / pressed | No change | None |
| Disabled | Not applicable | None |
| Loading | Not applicable | None |
| Error | Not applicable | None |
| Empty | Not applicable | None |

**Accessibility** — static text only, no role/ARIA, no keyboard interaction, minimum hit target not applicable.

## 3. Content and formatting

- Voice and tone in one line: plain, minimal, no marketing copy.
- Date, time, number, and currency formats: not used.
- Capitalization rule for buttons, headings, and labels: sentence case for visible text.
- Empty-state and error-message wording pattern: not used.

## 4. Known deviations

Places where the approved design does not follow its own rules or the
anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| Layout / typography | No explicit semantic components beyond one static text block | Product is one-screen proof-of-pipeline, so only one reusable pattern exists | None |
| Layout / motion | No hover, focus, active, disabled, loading, error, or empty states shown | Design is static by intent; extra states would be invented | None |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-02-14 | Initial design system for single-page proof-of-pipeline app | pending |
