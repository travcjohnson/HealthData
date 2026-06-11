/* Ridgeline — views and wiring. */
(() => {
  const $ = sel => document.querySelector(sel);
  const $$ = sel => Array.from(document.querySelectorAll(sel));

  const esc = s => String(s ?? "").replace(/[&<>"']/g, c => ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[c]));

  const fmtDay = d => new Date(d + "T12:00:00").toLocaleDateString("en-US", { weekday: "short", month: "short", day: "numeric" });
  const fmtMonth = m => new Date(m + "-15T12:00:00").toLocaleDateString("en-US", { month: "short", year: "2-digit" });

  Store.seedDemo();

  // ---------- routing ----------
  const TABS = ["dashboard", "habits", "training", "recovery", "labs", "okrs", "data"];
  function route() {
    const hash = (location.hash || "#dashboard").slice(1);
    const tab = TABS.includes(hash) ? hash : "dashboard";
    $$(".tab").forEach(b => b.classList.toggle("active", b.dataset.tab === tab));
    $$(".view").forEach(v => v.classList.toggle("active", v.id === "view-" + tab));
    render(tab);
  }
  window.addEventListener("hashchange", route);
  $$(".tab").forEach(b => b.addEventListener("click", () => (location.hash = b.dataset.tab)));

  function render(tab) {
    ({ dashboard: renderDashboard, habits: renderHabits, training: renderTraining,
       recovery: renderRecovery, labs: renderLabs, okrs: renderOKRs, data: renderData }[tab])();
    $("#demo-banner").style.display = Store.hasDemo() ? "flex" : "none";
  }

  $("#clear-demo").addEventListener("click", () => {
    if (confirm("Remove all sample data? Your own logs and imports are kept.")) {
      Store.clearDemo();
      route();
    }
  });

  // ---------- dashboard ----------
  function renderDashboard() {
    const today = Store.todayStr();
    const focus = OKR.todayFocus();
    const w = Store.state.whoop[today] || Store.state.whoop[Store.dateOffset(today, -1)] || {};
    const week = OKR.scoreWeek(Store.weekStart(today));
    const month = OKR.scoreMonth(today.slice(0, 7));

    $("#dash-date").textContent = new Date().toLocaleDateString("en-US", { weekday: "long", month: "long", day: "numeric", year: "numeric" });

    const fc = $("#focus-card");
    fc.className = "card focus " + focus.zone;
    fc.innerHTML = `
      <div class="focus-head">${esc(focus.headline)}</div>
      <p>${esc(focus.detail)}</p>`;

    const race = Store.targetEvent();
    const days = OKR.daysBetween(today, race.date);
    $("#race-card").innerHTML = days >= 0
      ? `<div class="big-num">${days === 0 ? "🏁" : days}</div>
         <div><strong>${days === 0 ? "RACE DAY" : "days to"}</strong><br>${esc(race.name)}<br><span class="muted small">${esc(race.date)} · ${esc(race.kind)}</span></div>`
      : `<div><strong>${esc(race.name)}</strong> is done — set the next target on the Data tab.</div>`;

    $("#vitals").innerHTML = `
      <div class="vital"><span class="v-num ${w.recovery >= 67 ? "green" : w.recovery >= 34 ? "yellow" : w.recovery != null ? "red" : ""}">${w.recovery ?? "–"}</span><span class="v-lbl">Recovery %</span></div>
      <div class="vital"><span class="v-num">${w.sleepHours ?? "–"}</span><span class="v-lbl">Sleep h</span></div>
      <div class="vital"><span class="v-num">${w.strain ?? "–"}</span><span class="v-lbl">Strain</span></div>
      <div class="vital"><span class="v-num">${week.volume.hours}</span><span class="v-lbl">Week hrs</span></div>`;

    Charts.ring($("#ring-week"), week.score, "week");
    Charts.ring($("#ring-month"), month.score, "month");
    $("#dash-okr-summary").innerHTML = `
      <p><strong>${esc(week.objective)}</strong></p>
      ${week.krs.map(k => krRow(k)).join("")}`;

    // Today's habit quick-check
    const checks = Store.checksFor(today);
    $("#dash-habits").innerHTML = Store.activeHabits().map(h => `
      <button class="habit-chip ${checks[h.id] ? "done" : ""}" data-habit="${h.id}">
        <span>${h.emoji}</span> ${esc(h.name)}
      </button>`).join("");
    $$("#dash-habits .habit-chip").forEach(b =>
      b.addEventListener("click", () => { Store.toggleCheck(today, b.dataset.habit); renderDashboard(); }));

    const trend = last14Recovery();
    Charts.line($("#dash-recovery-chart"), trend, {
      min: 0, max: 100,
      zones: [{ from: 0, to: 33, color: "#e54848" }, { from: 33, to: 66, color: "#e5c048" }, { from: 66, to: 100, color: "#3ec97e" }],
    });
  }

  function last14Recovery() {
    const today = Store.todayStr();
    const out = [];
    for (let i = 13; i >= 0; i--) {
      const d = Store.dateOffset(today, -i);
      const w = Store.state.whoop[d];
      out.push({ label: d.slice(5).replace("-", "/"), value: w ? w.recovery : null });
    }
    return out;
  }

  function krRow(k) {
    const pct = Math.round(k.score * 100);
    const col = k.score >= 0.7 ? "green" : k.score >= 0.4 ? "yellow" : "red";
    return `<div class="kr">
      <div class="kr-top"><span>${esc(k.name)}</span><span class="muted">${esc(k.actual)}</span></div>
      <div class="bar-bg"><div class="bar-fill ${col}" style="width:${pct}%"></div></div>
    </div>`;
  }

  // ---------- habits ----------
  function renderHabits() {
    const today = Store.todayStr();
    const ws = Store.weekStart(today);
    const days = [...Array(7)].map((_, i) => Store.dateOffset(ws, i));
    const habits = Store.activeHabits();

    let html = `<table class="habit-grid"><thead><tr><th>Habit</th>`;
    for (const d of days) {
      const isToday = d === today;
      html += `<th class="${isToday ? "today-col" : ""}">${["Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"][days.indexOf(d)]}<br><span class="muted small">${d.slice(8)}</span></th>`;
    }
    html += `<th>Streak</th><th></th></tr></thead><tbody>`;
    for (const h of habits) {
      html += `<tr><td class="habit-name">${h.emoji} ${esc(h.name)}${h.why ? `<div class="muted small">${esc(h.why)}</div>` : ""}</td>`;
      for (const d of days) {
        const done = !!(Store.checksFor(d)[h.id]);
        const future = d > today;
        html += `<td>${future
          ? `<span class="cell future"></span>`
          : `<button class="cell ${done ? "done" : ""}" data-d="${d}" data-h="${h.id}" aria-label="toggle">${done ? "✓" : ""}</button>`}</td>`;
      }
      const streak = Store.habitStreak(h.id);
      html += `<td class="streak">${streak ? "🔥 " + streak : "–"}</td>`;
      html += `<td><button class="link-btn" data-archive="${h.id}">archive</button></td></tr>`;
    }
    html += `</tbody></table>`;

    const adh = Store.habitAdherence(ws, today);
    const t = OKR.weeklyTargets(ws);
    html += `<p class="muted">Week so far: <strong>${Math.round(adh.pct * 100)}%</strong> adherence (${adh.done}/${adh.possible}) · this week's bar: <strong>${t.habitTargetPct}%</strong></p>`;

    $("#habit-table").innerHTML = html;
    $$("#habit-table .cell[data-d]").forEach(b =>
      b.addEventListener("click", () => { Store.toggleCheck(b.dataset.d, b.dataset.h); renderHabits(); }));
    $$("#habit-table [data-archive]").forEach(b =>
      b.addEventListener("click", () => {
        if (confirm("Archive this habit? Past checks are kept.")) { Store.archiveHabit(b.dataset.archive); renderHabits(); }
      }));
  }

  $("#habit-add-form").addEventListener("submit", e => {
    e.preventDefault();
    const name = $("#habit-new-name").value.trim();
    if (!name) return;
    Store.addHabit(name, $("#habit-new-emoji").value.trim());
    $("#habit-new-name").value = ""; $("#habit-new-emoji").value = "";
    renderHabits();
  });

  // ---------- training ----------
  function renderTraining() {
    // Real Strava history + in-app months appended.
    const histData = SEED.stravaMonthly.map(m => ({
      label: fmtMonth(m.month), value: m.hours, faded: !!m.estimated, accent: !!m.pr,
    }));
    const loggedMonths = {};
    for (const a of Store.state.activities) {
      const k = a.date.slice(0, 7);
      loggedMonths[k] = (loggedMonths[k] || 0) + (a.hours || 0);
    }
    for (const k of Object.keys(loggedMonths).sort()) {
      if (!SEED.stravaMonthly.some(m => m.month === k)) {
        histData.push({ label: fmtMonth(k), value: +loggedMonths[k].toFixed(1), accent: true });
      }
    }
    Charts.bars($("#training-history-chart"), histData, { height: 200, values: false });

    const pr = SEED.stravaMonthly.find(m => m.pr);
    $("#training-history-note").innerHTML =
      `Real history from your Strava recaps. <strong>May 2024 PR month:</strong> ${pr.miles.toLocaleString()} mi · ${pr.hours}h · ${pr.elevFt.toLocaleString()} ft over ${pr.activeDays} active days. Lighter bars are estimates from recap charts; orange bars are PR/current.`;

    // Last 8 weeks volume
    const ws = Store.weekStart(Store.todayStr());
    const weeks = [...Array(8)].map((_, i) => Store.dateOffset(ws, -7 * (7 - i)));
    Charts.bars($("#training-weeks-chart"), weeks.map(w => ({
      label: w.slice(5).replace("-", "/"),
      value: Store.weekVolume(w).hours,
      accent: w === ws,
    })), { height: 150, values: true, fmt: v => v + "h" });

    const recent = [...Store.state.activities].sort((a, b) => b.date.localeCompare(a.date)).slice(0, 12);
    $("#activity-list").innerHTML = recent.length ? recent.map(a => `
      <div class="row-item">
        <div>
          <strong>${esc(a.name || a.type)}</strong> <span class="pill">${esc(a.type)}</span>
          ${a.demo ? `<span class="pill demo">sample</span>` : ""}
          <div class="muted small">${fmtDay(a.date)} · ${a.hours}h · ${a.miles} mi · ${(a.elevFt || 0).toLocaleString()} ft</div>
        </div>
        <button class="link-btn" data-del="${a.id}">delete</button>
      </div>`).join("") : `<p class="muted">No activities yet — log one above or import on the Data tab.</p>`;
    $$("#activity-list [data-del]").forEach(b =>
      b.addEventListener("click", () => { Store.deleteActivity(b.dataset.del); renderTraining(); }));
  }

  $("#activity-form").addEventListener("submit", e => {
    e.preventDefault();
    const f = e.target;
    Store.logActivity({
      date: f.elements.date.value || Store.todayStr(),
      type: f.elements.type.value,
      name: f.elements.name.value.trim() || f.elements.type.value,
      hours: +f.elements.hours.value || 0,
      miles: +f.elements.miles.value || 0,
      elevFt: +f.elements.elev.value || 0,
    });
    f.reset();
    renderTraining();
  });

  // ---------- recovery ----------
  function renderRecovery() {
    const today = Store.todayStr();
    const rows = [];
    for (let i = 27; i >= 0; i--) {
      const d = Store.dateOffset(today, -i);
      const w = Store.state.whoop[d];
      rows.push({ label: d.slice(5).replace("-", "/"), date: d, ...(w || {}) });
    }
    Charts.line($("#recovery-chart"), rows.map(r => ({ label: r.label, value: r.recovery ?? null })), {
      min: 0, max: 100,
      zones: [{ from: 0, to: 33, color: "#e54848" }, { from: 33, to: 66, color: "#e5c048" }, { from: 66, to: 100, color: "#3ec97e" }],
    });
    Charts.line($("#sleep-chart"), rows.map(r => ({ label: r.label, value: Store.dailySleep(r.date) })), { min: 4, max: 10 });

    // Apple Health panel appears once an export has been imported.
    const apple = Store.state.apple || {};
    const appleRows = rows.map(r => ({ label: r.label, ...(apple[r.date] || {}) }));
    const hasApple = appleRows.some(r => r.hrv != null || r.rhr != null);
    $("#apple-card").hidden = !hasApple;
    if (hasApple) {
      Charts.line($("#hrv-chart"), appleRows.map(r => ({ label: r.label, value: r.hrv ?? null })), {});
      Charts.line($("#rhr-chart"), appleRows.map(r => ({ label: r.label, value: r.rhr ?? null })), {});
    }

    const week = Store.whoopBetween(Store.weekStart(today), today);
    const sleep7 = week.filter(w => w.sleepHours >= 7).length;
    const avgRec = Store.avgRecovery(Store.dateOffset(today, -6), today);
    $("#recovery-stats").innerHTML = `
      <div class="vital"><span class="v-num">${avgRec ?? "–"}</span><span class="v-lbl">7-day avg recovery</span></div>
      <div class="vital"><span class="v-num">${sleep7}</span><span class="v-lbl">≥7h nights this week</span></div>`;

    const f = $("#whoop-form");
    f.elements.date.value = today;
  }

  $("#whoop-form").addEventListener("submit", e => {
    e.preventDefault();
    const f = e.target;
    Store.logWhoop(f.elements.date.value, f.elements.recovery.value === "" ? null : +f.elements.recovery.value,
      f.elements.sleep.value === "" ? null : +f.elements.sleep.value, f.elements.strain.value === "" ? null : +f.elements.strain.value);
    renderRecovery();
  });

  // ---------- labs ----------
  const LAB_FLAGS = { watch: ["yellow", "watch"], action: ["red", "act now"], improving: ["green", "improving"], info: ["", "context"] };

  function renderLabs() {
    const personal = Store.state.labs;

    $("#labs-source").textContent = personal
      ? "Recovered " + (personal.recoveredAt || "") + " · " + (personal.source || "")
      : "No private lab record loaded in this browser yet.";

    $("#labs-insights").innerHTML = personal && personal.insights.length
      ? personal.insights.map(i => {
          const [col, label] = LAB_FLAGS[i.flag] || ["", ""];
          return `<div class="row-item">
            <div><strong>${esc(i.marker)}</strong>
              <div class="muted small">${esc(i.text)}</div>
            </div>
            <span class="score-pill ${col}">${label}</span>
          </div>`;
        }).join("")
      : `<p class="muted">Your lab record stays out of the public site by design. Import your private
         <strong>labs JSON</strong> on the Data tab to light this up in this browser.</p>`;

    const values = personal ? personal.values : [];
    const latest = {};
    for (const v of [...values].sort((a, b) => a.date.localeCompare(b.date))) latest[v.slug] = v;

    let html = `<table class="habit-grid"><thead><tr>
      <th style="text-align:left">Marker</th><th style="text-align:left">Category</th>
      <th>Optimal</th><th>Latest known</th><th style="text-align:left">Why it matters</th></tr></thead><tbody>`;
    for (const m of LABS_CATALOG) {
      const v = latest[m.slug];
      const history = values.filter(x => x.slug === m.slug).sort((a, b) => a.date.localeCompare(b.date));
      const trend = history.length > 1 ? ` <span class="muted small">(${history.map(h => h.value).join(" → ")})</span>` : "";
      html += `<tr>
        <td class="habit-name"><strong>${esc(m.name)}</strong></td>
        <td style="text-align:left" class="muted small">${esc(m.cat)}</td>
        <td class="muted small">${esc(m.target)} ${esc(m.unit)}</td>
        <td>${v ? `<strong>${v.value}</strong>${v.approx ? "≈" : ""}${trend}<div class="muted small">${esc(v.date)}</div>` : `<span class="muted small">–</span>`}</td>
        <td style="text-align:left" class="muted small">${esc(m.why)}</td>
      </tr>`;
    }
    html += `</tbody></table>`;
    $("#labs-table").innerHTML = html;

    $("#labs-reports").innerHTML = personal && personal.reports.length
      ? personal.reports.map(r => `
        <div class="row-item">
          <div><strong>${fmtDay(r.date)}</strong> <span class="pill">${esc(r.source)}</span>
            <div class="muted small">${r.items.map(esc).join(" · ")}</div>
          </div>
        </div>`).join("")
      : `<p class="muted">Imports on the Data tab.</p>`;
  }

  // ---------- okrs ----------
  function renderOKRs() {
    const today = Store.todayStr();
    const ws = Store.weekStart(today);
    const week = OKR.scoreWeek(ws);
    const month = OKR.scoreMonth(today.slice(0, 7));

    // 8-week objective-score trend (current week is partial)
    const trend = [...Array(8)].map((_, i) => {
      const w = Store.dateOffset(ws, -7 * (7 - i));
      return { label: w.slice(5).replace("-", "/"), value: Math.round(OKR.scoreWeek(w).score * 100) };
    });
    Charts.line($("#okr-trend-chart"), trend, { min: 0, max: 100 });

    // month-over-month comparison
    const thisM = today.slice(0, 7);
    const [y, m] = thisM.split("-").map(Number);
    const prevM = new Date(y, m - 2, 15);
    const prevKey = prevM.getFullYear() + "-" + String(prevM.getMonth() + 1).padStart(2, "0");
    const a = OKR.monthStats(thisM), b = OKR.monthStats(prevKey);
    const rows = [
      ["Training hours", a.hours, b.hours, "h"],
      ["Distance", a.miles, b.miles, " mi"],
      ["Climbing", a.elevFt, b.elevFt, " ft"],
      ["Active days", a.activeDays, b.activeDays, ""],
      ["Avg recovery", a.avgRecovery, b.avgRecovery, "%"],
      ["≥7h sleep nights", a.sleepNights7, b.sleepNights7, ""],
      ["Habit adherence", a.habitPct, b.habitPct, "%"],
    ];
    $("#month-compare").innerHTML = rows.map(([name, cur, prev, unit]) => {
      const has = cur != null && prev != null;
      const delta = has ? cur - prev : null;
      const col = !has || delta === 0 ? "muted" : delta > 0 ? "green" : "red";
      const arrow = !has ? "" : delta > 0 ? "▲" : delta < 0 ? "▼" : "–";
      return `<div class="kr-top" style="padding:5px 0">
        <span>${esc(name)}</span>
        <span><strong>${cur ?? "–"}${cur != null ? unit : ""}</strong>
          <span class="muted small">vs ${prev ?? "–"}${prev != null ? unit : ""}</span>
          <span class="${col}">${arrow}${has && delta !== 0 ? Math.abs(Math.round(delta * 10) / 10) : ""}</span></span>
      </div>`;
    }).join("");

    $("#okr-week").innerHTML = okrCard(
      `This week · ${fmtDay(ws)} → ${fmtDay(Store.dateOffset(ws, 6))}`,
      week.objective, week.krs, week.score, week.targets.reasons,
      `Phase: <strong>${week.targets.phase}</strong>${week.targets.recAvg != null ? ` · trailing recovery ${week.targets.recAvg}%` : ""} · baseline ${week.targets.baselineHours}h/wk`
    );

    $("#okr-month").innerHTML = okrCard(
      `This month · ${new Date().toLocaleDateString("en-US", { month: "long", year: "numeric" })}`,
      month.objective, month.krs, month.score, month.targets.reasons,
      `So far: ${month.stats.miles} mi · ${month.stats.hours}h · ${month.stats.elevFt.toLocaleString()} ft · ${month.stats.activeDays} active days`
    );

    // Past 4 weeks review
    const past = [1, 2, 3, 4].map(i => OKR.scoreWeek(Store.dateOffset(ws, -7 * i)));
    $("#okr-history").innerHTML = past.map(p => {
      const pct = Math.round(p.score * 100);
      const col = p.score >= 0.7 ? "green" : p.score >= 0.4 ? "yellow" : "red";
      return `<div class="row-item">
        <div><strong>Week of ${fmtDay(p.weekStart)}</strong>
          <div class="muted small">${p.volume.hours}h · ${p.volume.activeDays} active days · ${esc(p.objective)}</div>
        </div>
        <span class="score-pill ${col}">${pct}%</span>
      </div>`;
    }).join("");
  }

  function okrCard(period, objective, krs, score, reasons, meta) {
    const pct = Math.round(score * 100);
    const col = score >= 0.7 ? "green" : score >= 0.4 ? "yellow" : "red";
    return `
      <div class="okr-head">
        <div>
          <div class="muted small">${esc(period)}</div>
          <h3>${esc(objective)}</h3>
          <div class="muted small">${meta}</div>
        </div>
        <span class="score-pill big ${col}">${pct}%</span>
      </div>
      ${krs.map(krRow).join("")}
      <details class="why">
        <summary>Why these targets?</summary>
        <ul>${reasons.map(r => `<li>${esc(r)}</li>`).join("")}</ul>
      </details>`;
  }

  // ---------- data ----------
  function renderData() {
    const ev = Store.targetEvent();
    $("#profile-box").innerHTML = `
      <p><strong>${esc(SEED.profile.name)}</strong> · ${esc(SEED.profile.location)}<br>
      ${esc(SEED.profile.primarySport)} · Strava athlete #${SEED.profile.stravaAthleteId}<br>
      <span class="muted small">Target: ${esc(ev.name)} — ${esc(ev.date)}</span></p>`;

    const counts = {
      activities: Store.state.activities.length,
      whoopDays: Object.keys(Store.state.whoop).length,
      appleDays: Object.keys(Store.state.apple || {}).length,
      habitChecks: Object.values(Store.state.checks).reduce((s, c) => s + Object.keys(c).length, 0),
    };
    $("#data-counts").textContent =
      `${counts.activities} activities · ${counts.whoopDays} recovery days · ${counts.appleDays} Apple Health days · ${counts.habitChecks} habit checks stored locally in this browser.`;

    const f = $("#target-form");
    f.elements.tname.value = ev.name;
    f.elements.tdate.value = ev.date;
    f.elements.tkind.value = ev.kind || "";
  }

  $("#target-form").addEventListener("submit", e => {
    e.preventDefault();
    const f = e.target;
    Store.setTargetEvent({
      name: f.elements.tname.value.trim(),
      date: f.elements.tdate.value,
      kind: f.elements.tkind.value.trim(),
    });
    renderData();
    alert("Target updated — the OKR engine now plans around it.");
  });

  $("#target-reset").addEventListener("click", () => {
    Store.setTargetEvent(null);
    renderData();
  });

  $("#import-apple").addEventListener("change", async e => {
    const file = e.target.files[0];
    if (!file) return;
    const prog = $("#apple-progress");
    prog.textContent = "Scanning… 0%";
    try {
      const res = await Store.importAppleHealthXML(file, p => {
        prog.textContent = "Scanning… " + Math.round(p * 100) + "%";
      });
      prog.textContent = `Done: ${res.days} days of metrics, ${res.workouts} workouts, ${res.records.toLocaleString()} records.`;
      route();
    } catch (err) {
      prog.textContent = "Import failed: " + err.message;
    }
    e.target.value = "";
  });

  $("#export-btn").addEventListener("click", () => {
    const blob = new Blob([Store.exportJSON()], { type: "application/json" });
    const a = document.createElement("a");
    a.href = URL.createObjectURL(blob);
    a.download = "ridgeline-export-" + Store.todayStr() + ".json";
    a.click();
  });

  function fileHandler(inputId, fn) {
    $(inputId).addEventListener("change", e => {
      const file = e.target.files[0];
      if (!file) return;
      const reader = new FileReader();
      reader.onload = () => {
        try {
          const result = fn(reader.result);
          alert(typeof result === "number" ? `Imported ${result} rows.` : "Import complete.");
          route();
        } catch (err) {
          alert("Import failed: " + err.message);
        }
        e.target.value = "";
      };
      reader.readAsText(file);
    });
  }
  fileHandler("#import-json", t => Store.importJSON(t));
  fileHandler("#import-whoop", t => Store.importWhoopCSV(t));
  fileHandler("#import-strava", t => Store.importStravaCSV(t));
  fileHandler("#import-labs", t => Store.importLabsJSON(t));

  if ("serviceWorker" in navigator) {
    navigator.serviceWorker.register("sw.js").catch(() => { /* http/file context — fine */ });
  }

  route();
})();
