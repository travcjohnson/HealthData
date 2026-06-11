# Ridgeline — Travis's personal health OS

A deeply personal health tracker that combines **Whoop recovery**, **Strava training**,
and a **daily habit tracker** into auto-generated **weekly & monthly OKRs** that adapt
to what actually happened.

Built as a zero-dependency static web app: open `index.html`, or deploy the `webapp/`
folder to any static host (Vercel, Netlify, GitHub Pages).

## What's personal about it

The seed data is reconstructed from Travis's real Strava monthly recap emails:

| Month | Hours | Miles | Elevation | Active days |
|-------|------:|------:|----------:|------------:|
| Apr 2024 | 64 | 953 | 57,772 ft | 22 |
| **May 2024 (PR)** | **73** | **1,131** | **61,499 ft** | **25** |
| Jun 2024 | 32 | 549 | 29,856 ft | — |
| Jul 2024 | 30 | 484 | 31,742 ft | 12 |
| Aug 2024 | 34 | 520 | 31,863 ft | — |
| Sep 2024 | 15 | 202 | 11,345 ft | 7 |

(Jun 2023 – Mar 2024 are estimated from the recap bar charts; the scale was verified
against the months with actual numbers.)

Target event: **Cuyamaca Epic Team Race — June 20, 2026** (from the Strava club emails).

## The OKR engine

Targets are computed, not hand-set, and each card shows *why*:

- **Progressive overload**: weekly hours target = 3-week rolling baseline × 1.10
- **Automatic taper**: ≤14 days to race → 70% of baseline; race week → 50%, and
  *overshooting* the taper target scores **down**
- **Recovery-aware**: trailing 7-day Whoop recovery < 50% cuts the volume target 30%,
  50–65% cuts it 15%
- **Habit ratchet**: this week's adherence bar = last week's actual + 10 points
  (capped at 90% — always reachable, never comfortable)

## Data in / data out

- **Recovery tab** — 30-second morning log (recovery %, sleep, strain), or import the
  Whoop data export (`physiological_cycles.csv`)
- **Training tab** — quick activity log, or import the Strava bulk export
  (`activities.csv`)
- **Habits tab** — Monday-start week grid, streaks, add/archive habits
- **Data tab** — full JSON export/import

Everything is stored in the browser's localStorage; nothing leaves the device.
First load seeds 28 days of clearly-labeled sample recovery/activity data so the app
demonstrates itself — one click clears it.

## Run locally

```bash
cd webapp && python3 -m http.server 8080   # or just open index.html
```
