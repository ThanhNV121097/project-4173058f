# Design System — hello-word-28

> Source of truth: the approved `index.html` (preview: approved design).
> Every value below is extracted from it. Changing a value here without
> changing the approved design is a defect.

Last updated: 2025-02-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#ffffff` | Page background |
| `--color-surface` | `#ffffff` | Card / panel background |
| `--color-surface-raised` | `#ffffff` | Modal, dropdown, popover |
| `--color-border` | `#e5e5e5` | Default border, divider |
| `--color-text` | `#000000` | Body text |
| `--color-text-muted` | `#666666` | Secondary text, captions |
| `--color-primary` | `#000000` | Primary action background |
| `--color-primary-text` | `#ffffff` | Text on primary |
| `--color-success` | `#000000` | Success state |
| `--color-warning` | `#666666` | Warning state |
| `--color-danger` | `#000000` | Destructive action, error |
| `--color-focus` | `#000000` | Focus ring |

#### Contrast audit

Every text-on-background pair actually used. Body text ≥ 4.5:1, large text (≥ 18.66px bold or ≥ 24px) ≥ 3:1, UI borders ≥ 3:1.

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |
| `--color-text-muted` | `--color-bg` | `5.74:1` | AA / AA Large |
| `--color-primary-text` | `--color-primary` | `21:1` | AA / AA Large |
| `--color-text-muted` | `--color-surface` | `5.74:1` | AA / AA Large |

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

- Body: `Arial, Helvetica, sans-serif` via system fallback; no webfont.
- Headings: `Arial, Helvetica, sans-serif` via system fallback; no webfont.
- Mono: not used.

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-xs` | `0.875rem` | `1.2` | `400` | Captions, helper text, labels |
| `--text-sm` | `1rem` | `1.2` | `400` | Secondary body |
| `--text-base` | `1.125rem` | `1.2` | `400` | Body |
| `--text-lg` | `1.125rem` | `1.2` | `400` | Lead paragraph |
| `--text-xl` | `2rem` to `4.5rem` fluid | `1.1` | `400` | h1 |
| `--text-2xl` | not used | | | h2 |
| `--text-3xl` | not used | | | h1 larger range |

Heading levels are used in order and never skipped for visual sizing.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--radius-sm` | `16px` | Input, badge |
| `--radius-md` | `20px` | Button, card |
| `--radius-lg` | `20px` | Modal |
| `--radius-full` | `9999px` | Avatar, pill |
| `--border-width` | `1px` | Default border |
| `--shadow-sm` | `0 1px 0 rgba(0,0,0,0.02)` | Resting card |
| `--shadow-md` | not used | Dropdown, popover |
| `--shadow-lg` | not used | Modal |
| `--duration-fast` | not used | Hover, focus |
| `--duration-base` | not used | Panel open/close |
| `--easing` | not used | All transitions |

Motion respects `prefers-reduced-motion: reduce`: no motion exists, so state changes remain static.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | `0px` | fluid | 1 | `24px` |
| `md` | `641px` | fluid | 1 | `24px` |
| `lg` | not used | | | |
| `xl` | not used | | | |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | not used |
| Dropdown | not used |
| Modal backdrop | not used |
| Modal | not used |
| Toast | not used |

## 2. Components

One subsection per reusable component. Every component lists **all** states.

### 2.1 Greeting frame

**Purpose** — Centered white panel for greeting states. Use for loaded, loading, and error. Do not use for navigation, forms, or multi-panel layouts.

**Anatomy** — `[label] [content block]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default panel | `--color-surface`, `--color-border`, `--radius-md`, `--shadow-sm` | Main content frame |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | `min-height: 220px` | `24px` outer section padding | `--text-xl` / `--text-base` |

**States** — every row must be filled in.

| State | Visual change | Tokens |
|---|---|---|
| Default | White panel with centered greeting | `--color-surface`, `--color-text` |
| Hover | No change | none |
| Focus (keyboard) | Not interactive | none |
| Active / pressed | Not interactive | none |
| Disabled | Not interactive | none |
| Loading | Shows muted loading note below greeting | `--color-text-muted` |
| Error | Shows muted error note below greeting | `--color-text-muted` |
| Empty | Not used; product always has greeting content | none |

**Accessibility** — static content only; no role, no keyboard interaction. Minimum hit target not applicable.

### 2.2 State label

**Purpose** — Small uppercase status label at top-left of frame. Use only for loaded/loading/error markers.

**Anatomy** — `[text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-text-muted`, letter spacing `0.12em` | State marker |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | auto | none | `--text-xs` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Muted uppercase text | `--color-text-muted` |
| Hover | No change | none |
| Focus (keyboard) | Not interactive | none |
| Active / pressed | Not interactive | none |
| Disabled | Not interactive | none |
| Loading | Same text style | none |
| Error | Same text style | none |
| Empty | Not used | none |

**Accessibility** — decorative label only; no role or keyboard handling.

### 2.3 Greeting text

**Purpose** — Main greeting line. Use as single large line centered in frame.

**Anatomy** — `[text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-text`, `--text-xl`, `font-weight: 400`, `line-height: 1.1`, `letter-spacing: -0.04em` | Primary greeting |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Fluid | `clamp(2rem, 6vw, 4.5rem)` | none | `--text-xl` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Large centered black text | `--color-text` |
| Hover | No change | none |
| Focus (keyboard) | Not interactive | none |
| Active / pressed | Not interactive | none |
| Disabled | Not interactive | none |
| Loading | Same greeting text with loading note below | `--color-text-muted` |
| Error | Same greeting text with error note below | `--color-text-muted` |
| Empty | Not used | none |

**Accessibility** — rendered as heading in loaded state. No custom interaction.

### 2.4 Status note

**Purpose** — Helper line below greeting for loading and error states only.

**Anatomy** — `[text]`.

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Loading note | `--color-text-muted`, `--text-xs` | While data loads |
| Error note | `--color-text-muted`, `--text-xs` | When greeting read fails |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | auto | `10px` top margin from greeting | `--text-xs` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Muted helper copy | `--color-text-muted` |
| Hover | No change | none |
| Focus (keyboard) | Not interactive | none |
| Active / pressed | Not interactive | none |
| Disabled | Not interactive | none |
| Loading | Shows `Loading stored greeting…` | `--color-text-muted` |
| Error | Shows `Could not read greeting row.` | `--color-text-muted` |
| Empty | Not used | none |

**Accessibility** — static helper text; no role or keyboard interaction.

## 3. Content and formatting

- Voice and tone: plain, direct, no marketing copy.
- Date, time, number, and currency formats: not used in approved design.
- Capitalization rule: state labels use uppercase; greeting sentence case as written (`Hello Word`).
- Empty-state and error-message wording pattern: short sentence, plain noun phrase, no extra punctuation unless shown in approved copy.

## 4. Known deviations

Places where the approved design does not follow its own rules or the
anti-patterns in `references/ai-defaults.md`. Record, do not silently fix.

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| `Typography` | `--text-2xl` and `--text-3xl` are not used, but template expects a full ramp | Approved mockup only needs one fluid headline size | Keep unused tokens out of implementation unless later screens need them |
| `Radius` | `--radius-sm` and `--radius-lg` both map to `20px` in the source UI shape | Only one card shape exists in approved design | Leave as documented; no visible conflict in current screen |
| `Layout` | `lg` and `xl` breakpoints not used | Single-screen design never expands into multi-column layout | Add only if future screens need them |
| `Motion` | No transitions or easing in approved design | Static UI by choice | None |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-02-14 | Initial design system extracted from approved mockup | pending |
