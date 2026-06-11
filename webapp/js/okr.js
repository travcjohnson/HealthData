/* Ridgeline OKR engine.
 *
 * Generates weekly and monthly objectives from the data itself:
 *  - progressive overload capped at +10% per week off a 3-week rolling baseline
 *  - automatic taper inside 14 days of the target event
 *  - volume targets cut when average recovery trends low
 *  - habit targets ratchet up from last week's actual adherence
 * Every target comes with the reasoning that produced it.
 */
var OKR = (() => {
  const clamp = (v, lo, hi) => Math.min(hi, Math.max(lo, v));

  function daysBetween(a, b) {
    return Math.round((new Date(b + "T12:00:00") - new Date(a + "T12:00:00")) / 86400000);
  }

  function weeklyTargets(weekStartDate) {
    const S = Store;
    const reasons = [];

    // Baseline: average of the previous 3 weeks' training hours.
    const prevWeeks = [1, 2, 3].map(i => S.weekVolume(S.dateOffset(weekStartDate, -7 * i)));
    const baselineHours = prevWeeks.reduce((s, w) => s + w.hours, 0) / 3;

    const race = SEED.targetEvent;
    const daysToRace = daysBetween(weekStartDate, race.date);

    let phase = "build";
    let targetHours;
    if (daysToRace >= 0 && daysToRace <= 6) {
      phase = "race";
      targetHours = clamp(baselineHours * 0.5, 2, 8);
      reasons.push(`${race.name} is THIS WEEK (${race.date}). Volume cut to ~50% of baseline — show up fresh, not fit-on-paper.`);
    } else if (daysToRace > 6 && daysToRace <= 14) {
      phase = "taper";
      targetHours = clamp(baselineHours * 0.7, 3, 10);
      reasons.push(`${daysToRace} days to ${race.name}: taper week. Target trimmed to ~70% of your 3-week baseline (${baselineHours.toFixed(1)}h). Keep intensity, drop duration.`);
    } else {
      targetHours = clamp(baselineHours * 1.1, 4, 16);
      reasons.push(`Build phase: +10% on your 3-week baseline of ${baselineHours.toFixed(1)}h/week.`);
    }

    // Recovery modulation off the trailing 7 days before the week starts.
    const recAvg = S.avgRecovery(S.dateOffset(weekStartDate, -7), S.dateOffset(weekStartDate, -1));
    if (recAvg != null && recAvg < 50) {
      targetHours = +(targetHours * 0.7).toFixed(1);
      reasons.push(`Average recovery was ${recAvg}% (red zone) — volume target reduced 30%. The plan adapts; you don't dig the hole deeper.`);
    } else if (recAvg != null && recAvg < 65) {
      targetHours = +(targetHours * 0.85).toFixed(1);
      reasons.push(`Average recovery was ${recAvg}% (yellow) — volume target eased 15%.`);
    } else if (recAvg != null) {
      reasons.push(`Average recovery ${recAvg}% — green light to absorb the full load.`);
    } else {
      reasons.push(`No recovery data yet — log or import Whoop data to make this smarter.`);
    }
    targetHours = +targetHours.toFixed(1);

    // Habit target ratchets from last week's actual.
    const lastWeekStart = S.dateOffset(weekStartDate, -7);
    const lastAdh = S.habitAdherence(lastWeekStart, S.dateOffset(lastWeekStart, 6));
    const habitTargetPct = lastAdh.possible
      ? clamp(Math.round(lastAdh.pct * 100) + 10, 60, 90)
      : 70;
    if (lastAdh.possible) {
      reasons.push(`Habits last week: ${Math.round(lastAdh.pct * 100)}%. This week's bar: ${habitTargetPct}% — always reachable, never comfortable.`);
    }

    const sleepNightsTarget = phase === "build" ? 5 : 6;
    const activeDaysTarget = phase === "race" ? 3 : phase === "taper" ? 4 : 4;

    return { phase, daysToRace, targetHours, baselineHours: +baselineHours.toFixed(1), recAvg, habitTargetPct, sleepNightsTarget, activeDaysTarget, reasons };
  }

  function scoreWeek(weekStartDate) {
    const S = Store;
    const t = weeklyTargets(weekStartDate);
    const end = S.dateOffset(weekStartDate, 6);
    const vol = S.weekVolume(weekStartDate);
    const whoop = S.whoopBetween(weekStartDate, end);
    const sleepNights = whoop.filter(w => w.sleepHours != null && w.sleepHours >= 7).length;
    const adh = S.habitAdherence(weekStartDate, end);

    const objectiveName = t.phase === "race"
      ? `Race week: arrive at ${SEED.targetEvent.name} fresh and sharp`
      : t.phase === "taper"
        ? `Taper: convert fitness into freshness for ${SEED.targetEvent.name}`
        : "Build the engine without breaking the rider";

    const krs = [
      {
        name: t.phase === "build" ? `Train ${t.targetHours}h` : `Train ≤ plan: ${t.targetHours}h (taper means LESS)`,
        actual: `${vol.hours}h`,
        // During taper, overshooting is the failure mode — score caps and decays past target.
        score: t.phase === "build"
          ? clamp(vol.hours / t.targetHours, 0, 1)
          : vol.hours <= t.targetHours ? clamp(vol.hours / (t.targetHours * 0.6), 0, 1) : clamp(2 - vol.hours / t.targetHours, 0, 1),
      },
      {
        name: `Active ${t.activeDaysTarget}+ days`,
        actual: `${vol.activeDays} days`,
        score: clamp(vol.activeDays / t.activeDaysTarget, 0, 1),
      },
      {
        name: `Sleep ≥7h on ${t.sleepNightsTarget} nights`,
        actual: `${sleepNights} nights`,
        score: clamp(sleepNights / t.sleepNightsTarget, 0, 1),
      },
      {
        name: `Habit adherence ≥ ${t.habitTargetPct}%`,
        actual: `${Math.round(adh.pct * 100)}%`,
        score: clamp((adh.pct * 100) / t.habitTargetPct, 0, 1),
      },
    ];

    const score = krs.reduce((s, k) => s + k.score, 0) / krs.length;
    return { weekStart: weekStartDate, objective: objectiveName, targets: t, krs, score, volume: vol };
  }

  function monthlyTargets(monthStr) {
    const S = Store;
    const reasons = [];
    const hist = SEED.stravaMonthly;
    const pr = hist.find(m => m.pr);

    // Baseline: trailing 3 logged months of in-app activity; falls back to history.
    const [y, m] = monthStr.split("-").map(Number);
    const recent = [];
    for (let i = 1; i <= 3; i++) {
      const d = new Date(y, m - 1 - i, 15);
      const key = d.toISOString().slice(0, 7);
      const acts = S.state.activities.filter(a => a.date.startsWith(key));
      if (acts.length) {
        recent.push({
          miles: acts.reduce((s, a) => s + (a.miles || 0), 0),
          hours: acts.reduce((s, a) => s + (a.hours || 0), 0),
          elevFt: acts.reduce((s, a) => s + (a.elevFt || 0), 0),
        });
      }
    }

    let targetMiles, targetHours, targetElev;
    if (recent.length) {
      const avg = k => recent.reduce((s, r) => s + r[k], 0) / recent.length;
      targetMiles = Math.round(avg("miles") * 1.1);
      targetHours = Math.round(avg("hours") * 1.1);
      targetElev = Math.round(avg("elevFt") * 1.1);
      reasons.push(`+10% on your trailing ${recent.length}-month average.`);
    } else {
      // No logged history: anchor to the real Strava record at a sustainable fraction.
      targetMiles = 400; targetHours = 28; targetElev = 25000;
      reasons.push(`Anchored to your real Strava history — sustainable June baseline (your May 2024 PR month was ${pr.miles} mi / ${pr.hours}h; we build back toward it, not from it).`);
    }

    // Race month: respect the taper, then rebuild.
    if (SEED.targetEvent.date.startsWith(monthStr)) {
      reasons.push(`${SEED.targetEvent.name} lands this month — the win is racing well, not padding totals.`);
    }

    return { targetMiles, targetHours, targetElev, reasons, pr };
  }

  function scoreMonth(monthStr) {
    const S = Store;
    const t = monthlyTargets(monthStr);
    const acts = S.state.activities.filter(a => a.date.startsWith(monthStr));
    const miles = +acts.reduce((s, a) => s + (a.miles || 0), 0).toFixed(0);
    const hours = +acts.reduce((s, a) => s + (a.hours || 0), 0).toFixed(1);
    const elevFt = Math.round(acts.reduce((s, a) => s + (a.elevFt || 0), 0));
    const activeDays = new Set(acts.map(a => a.date)).size;

    // Sleep consistency across the month.
    const daysInMonth = new Date(+monthStr.slice(0, 4), +monthStr.slice(5, 7), 0).getDate();
    const whoop = S.whoopBetween(monthStr + "-01", monthStr + "-" + String(daysInMonth).padStart(2, "0"));
    const sleepOk = whoop.filter(w => w.sleepHours >= 7).length;
    const sleepTarget = 20;

    const adh = S.habitAdherence(monthStr + "-01", monthStr + "-" + String(daysInMonth).padStart(2, "0"));

    const krs = [
      { name: `Ride ${t.targetMiles} mi`, actual: `${miles} mi`, score: clamp(miles / t.targetMiles, 0, 1) },
      { name: `${t.targetHours}h training time`, actual: `${hours}h`, score: clamp(hours / t.targetHours, 0, 1) },
      { name: `Climb ${t.targetElev.toLocaleString()} ft`, actual: `${elevFt.toLocaleString()} ft`, score: clamp(elevFt / t.targetElev, 0, 1) },
      { name: `≥7h sleep on ${sleepTarget} nights`, actual: `${sleepOk} nights`, score: clamp(sleepOk / sleepTarget, 0, 1) },
      { name: `Habit adherence ≥ 80%`, actual: `${Math.round(adh.pct * 100)}%`, score: clamp(adh.pct / 0.8, 0, 1) },
    ];
    const score = krs.reduce((s, k) => s + k.score, 0) / krs.length;

    return {
      month: monthStr,
      objective: "Be measurably healthier than last month — and prove it",
      targets: t, krs, score,
      stats: { miles, hours, elevFt, activeDays },
    };
  }

  /* Today's recommendation from latest recovery + phase. */
  function todayFocus() {
    const S = Store;
    const today = S.todayStr();
    const w = S.state.whoop[today] || S.state.whoop[S.dateOffset(today, -1)];
    const t = weeklyTargets(S.weekStart(today));
    const daysToRace = daysBetween(today, SEED.targetEvent.date);

    let headline, detail, zone;
    if (w && w.recovery != null) {
      if (w.recovery >= 67) {
        zone = "green";
        headline = "Green light — make it count";
        detail = t.phase === "build"
          ? "Recovery is high. Today is the day for intensity or the long climb."
          : "Recovery is high — but you're tapering. Short and sharp, then sit on your hands.";
      } else if (w.recovery >= 34) {
        zone = "yellow";
        headline = "Steady — zone 2 day";
        detail = "Moderate recovery. Endurance pace builds the base without digging a hole.";
      } else {
        zone = "red";
        headline = "Recovery day — that's the workout";
        detail = "Low recovery. Yoga, walk, mobility, early night. Adaptation happens at rest.";
      }
    } else {
      zone = "none";
      headline = "No recovery data for today";
      detail = "Log this morning's recovery below, or import your Whoop export on the Data tab.";
    }

    if (daysToRace >= 0 && daysToRace <= 14) {
      detail += ` ${daysToRace === 0 ? "RACE DAY." : daysToRace + " days to " + SEED.targetEvent.name + "."}`;
    }
    return { headline, detail, zone, daysToRace, phase: t.phase };
  }

  return { weeklyTargets, scoreWeek, monthlyTargets, scoreMonth, todayFocus, daysBetween };
})();
