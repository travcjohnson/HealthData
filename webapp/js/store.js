/* Ridgeline — state, persistence, and sample-data bootstrap.
 * Everything lives in localStorage under one key so export/import is trivial.
 */
var Store = (() => {
  const KEY = "ridgeline.v1";

  const todayStr = () => new Date().toISOString().slice(0, 10);

  function dateOffset(base, days) {
    const d = new Date(base + "T12:00:00");
    d.setDate(d.getDate() + days);
    return d.toISOString().slice(0, 10);
  }

  // Monday-start week key for a date string.
  function weekStart(dateStr) {
    const d = new Date(dateStr + "T12:00:00");
    const day = (d.getDay() + 6) % 7; // Mon=0
    d.setDate(d.getDate() - day);
    return d.toISOString().slice(0, 10);
  }

  function monthKey(dateStr) {
    return dateStr.slice(0, 7);
  }

  function blank() {
    return {
      version: 1,
      habits: SEED.defaultHabits.map(h => ({ ...h, archived: false })),
      checks: {},      // { "YYYY-MM-DD": { habitId: true } }
      whoop: {},       // { "YYYY-MM-DD": { recovery, sleepHours, strain, demo? } }
      activities: [],  // { id, date, type, hours, miles, elevFt, name, demo? }
      okrHistory: {},  // { weekStartDate: { computed snapshot } } — filled by OKR engine
      settings: { demoSeeded: false, demoCleared: false },
    };
  }

  let state = load();

  function load() {
    try {
      const raw = localStorage.getItem(KEY);
      if (raw) return JSON.parse(raw);
    } catch (e) { /* corrupted state falls through to fresh */ }
    return blank();
  }

  function save() {
    localStorage.setItem(KEY, JSON.stringify(state));
  }

  // Deterministic PRNG so the sample data is stable between loads.
  function mulberry32(a) {
    return function () {
      a |= 0; a = (a + 0x6D2B79F5) | 0;
      let t = Math.imul(a ^ (a >>> 15), 1 | a);
      t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t;
      return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
    };
  }

  /* Seed 28 days of clearly-labeled sample recovery + activity data so the app
   * demonstrates itself on first load. Real logs / imports replace it. */
  function seedDemo() {
    if (state.settings.demoSeeded || state.settings.demoCleared) return;
    const rnd = mulberry32(9337896);
    const today = todayStr();
    let recovery = 62;
    for (let i = 28; i >= 1; i--) {
      const date = dateOffset(today, -i);
      const dow = new Date(date + "T12:00:00").getDay(); // 0=Sun
      recovery = Math.max(20, Math.min(95, recovery + (rnd() - 0.48) * 22));
      const sleepHours = +(6.1 + rnd() * 2.4).toFixed(1);
      let strain = +(6 + rnd() * 6).toFixed(1);

      // Riding pattern: Social Saturday long ride, Sunday spin, Tue/Thu workouts.
      let act = null;
      if (dow === 6) act = { type: "Gravel Ride", hours: +(2.8 + rnd() * 1.4).toFixed(1), name: "Social SATURDAY" };
      else if (dow === 0 && rnd() > 0.35) act = { type: "Ride", hours: +(1.4 + rnd()).toFixed(1), name: "Sunday spin" };
      else if ((dow === 2 || dow === 4) && rnd() > 0.3) act = { type: "Ride", hours: +(1.0 + rnd() * 0.8).toFixed(1), name: dow === 2 ? "Tuesday intervals" : "Thursday tempo" };
      else if (dow === 3 && rnd() > 0.5) act = { type: "Yoga", hours: 0.5, name: "Recovery yoga" };

      if (act) {
        const mph = act.type === "Yoga" ? 0 : 13 + rnd() * 3;
        state.activities.push({
          id: "demo-" + date,
          date,
          type: act.type,
          name: act.name,
          hours: act.hours,
          miles: +(act.hours * mph).toFixed(1),
          elevFt: act.type === "Yoga" ? 0 : Math.round(act.hours * (700 + rnd() * 500)),
          demo: true,
        });
        strain = Math.min(20, strain + act.hours * 3);
      }

      state.whoop[date] = {
        recovery: Math.round(recovery),
        sleepHours,
        strain: +strain.toFixed(1),
        demo: true,
      };

      // Habit checks ~72% adherence so the OKR engine has something to chew on.
      const dayChecks = {};
      for (const h of state.habits) if (rnd() < 0.72) dayChecks[h.id] = true;
      if (act && act.type !== "Yoga") dayChecks["h-train"] = true;
      state.checks[date] = dayChecks;
      (state.settings.demoCheckDates = state.settings.demoCheckDates || []).push(date);
    }
    state.settings.demoSeeded = true;
    save();
  }

  function clearDemo() {
    state.activities = state.activities.filter(a => !a.demo);
    for (const d of Object.keys(state.whoop)) if (state.whoop[d].demo) delete state.whoop[d];
    for (const d of state.settings.demoCheckDates || []) delete state.checks[d];
    delete state.settings.demoCheckDates;
    state.settings.demoCleared = true;
    state.settings.demoSeeded = false;
    save();
  }

  function hasDemo() {
    return state.activities.some(a => a.demo) || Object.values(state.whoop).some(w => w.demo);
  }

  // ---- mutations ----
  function toggleCheck(date, habitId) {
    state.checks[date] = state.checks[date] || {};
    if (state.checks[date][habitId]) delete state.checks[date][habitId];
    else state.checks[date][habitId] = true;
    save();
  }

  function addHabit(name, emoji) {
    state.habits.push({ id: "h-" + Date.now(), name, emoji: emoji || "✅", why: "", archived: false });
    save();
  }

  function archiveHabit(id) {
    const h = state.habits.find(h => h.id === id);
    if (h) { h.archived = true; save(); }
  }

  function logWhoop(date, recovery, sleepHours, strain) {
    state.whoop[date] = {
      recovery: recovery == null ? null : Math.round(recovery),
      sleepHours: sleepHours == null ? null : +sleepHours,
      strain: strain == null ? null : +strain,
    };
    save();
  }

  function logActivity(a) {
    state.activities.push({ id: "a-" + Date.now(), ...a });
    state.activities.sort((x, y) => x.date.localeCompare(y.date));
    save();
  }

  function deleteActivity(id) {
    state.activities = state.activities.filter(a => a.id !== id);
    save();
  }

  // ---- queries ----
  function activeHabits() { return state.habits.filter(h => !h.archived); }

  function checksFor(date) { return state.checks[date] || {}; }

  function habitAdherence(startDate, endDate) {
    const habits = activeHabits();
    if (!habits.length) return { done: 0, possible: 0, pct: 0 };
    let done = 0, possible = 0;
    for (let d = startDate; d <= endDate; d = dateOffset(d, 1)) {
      const checks = state.checks[d] || {};
      for (const h of habits) {
        possible++;
        if (checks[h.id]) done++;
      }
    }
    return { done, possible, pct: possible ? done / possible : 0 };
  }

  function habitStreak(habitId) {
    let streak = 0;
    let d = todayStr();
    // today counts if checked; otherwise start from yesterday
    if (!(state.checks[d] && state.checks[d][habitId])) d = dateOffset(d, -1);
    while (state.checks[d] && state.checks[d][habitId]) {
      streak++;
      d = dateOffset(d, -1);
    }
    return streak;
  }

  function activitiesBetween(startDate, endDate) {
    return state.activities.filter(a => a.date >= startDate && a.date <= endDate);
  }

  function weekVolume(weekStartDate) {
    const end = dateOffset(weekStartDate, 6);
    const acts = activitiesBetween(weekStartDate, end);
    return {
      hours: +acts.reduce((s, a) => s + (a.hours || 0), 0).toFixed(1),
      miles: +acts.reduce((s, a) => s + (a.miles || 0), 0).toFixed(1),
      elevFt: Math.round(acts.reduce((s, a) => s + (a.elevFt || 0), 0)),
      activeDays: new Set(acts.map(a => a.date)).size,
    };
  }

  function whoopBetween(startDate, endDate) {
    const out = [];
    for (let d = startDate; d <= endDate; d = dateOffset(d, 1)) {
      if (state.whoop[d]) out.push({ date: d, ...state.whoop[d] });
    }
    return out;
  }

  function avgRecovery(startDate, endDate) {
    const rows = whoopBetween(startDate, endDate).filter(w => w.recovery != null);
    if (!rows.length) return null;
    return Math.round(rows.reduce((s, w) => s + w.recovery, 0) / rows.length);
  }

  // ---- import / export ----
  function exportJSON() {
    return JSON.stringify(state, null, 2);
  }

  function importJSON(text) {
    const parsed = JSON.parse(text);
    if (!parsed || parsed.version !== 1) throw new Error("Not a Ridgeline export");
    state = parsed;
    save();
  }

  /* Tolerant CSV import for Whoop physiological cycles export.
   * Matches columns by keyword: date/cycle start, recovery, asleep/sleep, strain. */
  function importWhoopCSV(text) {
    const rows = parseCSV(text);
    if (rows.length < 2) throw new Error("CSV appears empty");
    const header = rows[0].map(c => c.toLowerCase());
    const dateCol = header.findIndex(c => c.includes("cycle start") || c === "date" || c.includes("day"));
    const recCol = header.findIndex(c => c.includes("recovery"));
    const sleepCol = header.findIndex(c => c.includes("asleep") || (c.includes("sleep") && c.includes("dur")));
    const strainCol = header.findIndex(c => c.includes("strain"));
    if (dateCol < 0) throw new Error("No date column found");
    let count = 0;
    for (const row of rows.slice(1)) {
      const rawDate = (row[dateCol] || "").slice(0, 10).replace(/\//g, "-");
      const date = normalizeDate(rawDate);
      if (!date) continue;
      const rec = recCol >= 0 ? parseFloat(row[recCol]) : NaN;
      let sleep = sleepCol >= 0 ? parseFloat(row[sleepCol]) : NaN;
      if (!isNaN(sleep) && sleep > 24) sleep = sleep / 60; // minutes → hours
      const strain = strainCol >= 0 ? parseFloat(row[strainCol]) : NaN;
      state.whoop[date] = {
        recovery: isNaN(rec) ? null : Math.round(rec),
        sleepHours: isNaN(sleep) ? null : +sleep.toFixed(1),
        strain: isNaN(strain) ? null : +strain.toFixed(1),
      };
      count++;
    }
    save();
    return count;
  }

  /* Tolerant CSV import for Strava bulk-export activities.csv. */
  function importStravaCSV(text) {
    const rows = parseCSV(text);
    if (rows.length < 2) throw new Error("CSV appears empty");
    const header = rows[0].map(c => c.toLowerCase());
    const dateCol = header.findIndex(c => c.includes("activity date") || c === "date");
    const typeCol = header.findIndex(c => c.includes("activity type") || c === "type");
    const nameCol = header.findIndex(c => c.includes("activity name") || c === "name");
    const timeCol = header.findIndex(c => c.includes("elapsed time") || c.includes("moving time"));
    const distCol = header.findIndex(c => c === "distance" || c.includes("distance"));
    const elevCol = header.findIndex(c => c.includes("elevation gain"));
    if (dateCol < 0) throw new Error("No 'Activity Date' column found");
    let count = 0;
    for (const row of rows.slice(1)) {
      const date = normalizeDate(row[dateCol]);
      if (!date) continue;
      const secs = timeCol >= 0 ? parseFloat(row[timeCol]) : NaN;
      let dist = distCol >= 0 ? parseFloat(row[distCol]) : NaN; // Strava exports km or m
      if (!isNaN(dist) && dist > 1000) dist = dist / 1609.34;   // meters → miles
      else if (!isNaN(dist)) dist = dist * 0.621371;            // km → miles
      const elevM = elevCol >= 0 ? parseFloat(row[elevCol]) : NaN;
      state.activities.push({
        id: "s-" + date + "-" + count,
        date,
        type: typeCol >= 0 ? row[typeCol] : "Ride",
        name: nameCol >= 0 ? row[nameCol] : "Imported activity",
        hours: isNaN(secs) ? 0 : +(secs / 3600).toFixed(2),
        miles: isNaN(dist) ? 0 : +dist.toFixed(1),
        elevFt: isNaN(elevM) ? 0 : Math.round(elevM * 3.28084),
      });
      count++;
    }
    state.activities.sort((x, y) => x.date.localeCompare(y.date));
    save();
    return count;
  }

  function normalizeDate(raw) {
    if (!raw) return null;
    const d = new Date(raw);
    if (isNaN(d.getTime())) return null;
    return d.toISOString().slice(0, 10);
  }

  // Minimal quoted-field CSV parser.
  function parseCSV(text) {
    const rows = [];
    let row = [], field = "", inQuotes = false;
    for (let i = 0; i < text.length; i++) {
      const c = text[i];
      if (inQuotes) {
        if (c === '"' && text[i + 1] === '"') { field += '"'; i++; }
        else if (c === '"') inQuotes = false;
        else field += c;
      } else if (c === '"') inQuotes = true;
      else if (c === ",") { row.push(field); field = ""; }
      else if (c === "\n" || c === "\r") {
        if (field !== "" || row.length) { row.push(field); rows.push(row); row = []; field = ""; }
        if (c === "\r" && text[i + 1] === "\n") i++;
      } else field += c;
    }
    if (field !== "" || row.length) { row.push(field); rows.push(row); }
    return rows.filter(r => r.length > 1 || (r[0] && r[0].trim()));
  }

  return {
    get state() { return state; },
    save, todayStr, dateOffset, weekStart, monthKey,
    seedDemo, clearDemo, hasDemo,
    toggleCheck, addHabit, archiveHabit, logWhoop, logActivity, deleteActivity,
    activeHabits, checksFor, habitAdherence, habitStreak,
    activitiesBetween, weekVolume, whoopBetween, avgRecovery,
    exportJSON, importJSON, importWhoopCSV, importStravaCSV,
  };
})();
