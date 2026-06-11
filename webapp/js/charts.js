/* Ridgeline — tiny hand-rolled SVG charts. No dependencies. */
var Charts = (() => {
  const NS = "http://www.w3.org/2000/svg";

  function el(tag, attrs, parent) {
    const e = document.createElementNS(NS, tag);
    for (const k in attrs) e.setAttribute(k, attrs[k]);
    if (parent) parent.appendChild(e);
    return e;
  }

  /* Vertical bar chart. data: [{label, value, accent?, faded?}] */
  function bars(container, data, opts = {}) {
    container.innerHTML = "";
    if (!data.length) return;
    const W = opts.width || container.clientWidth || 600;
    const H = opts.height || 180;
    const pad = { t: 14, r: 6, b: 26, l: 6 };
    const svg = el("svg", { viewBox: `0 0 ${W} ${H}`, width: "100%", height: H, role: "img" }, container);
    const max = Math.max(...data.map(d => d.value), 1);
    const bw = (W - pad.l - pad.r) / data.length;

    data.forEach((d, i) => {
      const h = ((H - pad.t - pad.b) * d.value) / max;
      const x = pad.l + i * bw + bw * 0.15;
      const y = H - pad.b - h;
      el("rect", {
        x, y, width: bw * 0.7, height: Math.max(h, 1), rx: 3,
        fill: d.accent ? "var(--accent)" : "var(--bar)",
        opacity: d.faded ? 0.45 : 1,
      }, svg);
      if (opts.values && d.value) {
        el("text", { x: x + bw * 0.35, y: y - 4, "text-anchor": "middle", class: "chart-val" }, svg)
          .textContent = opts.fmt ? opts.fmt(d.value) : d.value;
      }
      if (d.label && (data.length <= 16 || i % 2 === 0)) {
        el("text", { x: x + bw * 0.35, y: H - 8, "text-anchor": "middle", class: "chart-lbl" }, svg)
          .textContent = d.label;
      }
    });
  }

  /* Line chart with optional colored zone bands. data: [{label, value}] */
  function line(container, data, opts = {}) {
    container.innerHTML = "";
    const points = data.filter(d => d.value != null);
    if (points.length < 2) {
      container.innerHTML = `<p class="muted small">Not enough data yet — log a few days to see the trend.</p>`;
      return;
    }
    const W = opts.width || container.clientWidth || 600;
    const H = opts.height || 160;
    const pad = { t: 10, r: 8, b: 22, l: 30 };
    const svg = el("svg", { viewBox: `0 0 ${W} ${H}`, width: "100%", height: H, role: "img" }, container);
    const min = opts.min != null ? opts.min : Math.min(...points.map(d => d.value));
    const max = opts.max != null ? opts.max : Math.max(...points.map(d => d.value));
    const x = i => pad.l + (i * (W - pad.l - pad.r)) / (data.length - 1);
    const y = v => pad.t + (H - pad.t - pad.b) * (1 - (v - min) / (max - min || 1));

    if (opts.zones) {
      for (const z of opts.zones) {
        el("rect", {
          x: pad.l, y: y(z.to), width: W - pad.l - pad.r, height: Math.max(y(z.from) - y(z.to), 0),
          fill: z.color, opacity: 0.10,
        }, svg);
      }
    }
    // y-axis ticks
    for (const tv of [min, (min + max) / 2, max]) {
      el("text", { x: pad.l - 6, y: y(tv) + 4, "text-anchor": "end", class: "chart-lbl" }, svg)
        .textContent = Math.round(tv);
    }

    let path = "";
    data.forEach((d, i) => {
      if (d.value == null) return;
      path += (path ? " L " : "M ") + x(i).toFixed(1) + " " + y(d.value).toFixed(1);
    });
    el("path", { d: path, fill: "none", stroke: "var(--accent)", "stroke-width": 2.5, "stroke-linejoin": "round" }, svg);

    data.forEach((d, i) => {
      if (d.value == null) return;
      el("circle", { cx: x(i), cy: y(d.value), r: 3, fill: "var(--accent)" }, svg);
      if (d.label && (data.length <= 10 || i % Math.ceil(data.length / 8) === 0)) {
        el("text", { x: x(i), y: H - 6, "text-anchor": "middle", class: "chart-lbl" }, svg).textContent = d.label;
      }
    });
  }

  /* Score ring 0..1 */
  function ring(container, score, label) {
    container.innerHTML = "";
    const size = 84, r = 34, c = 2 * Math.PI * r;
    const svg = el("svg", { viewBox: `0 0 ${size} ${size}`, width: size, height: size }, container);
    el("circle", { cx: size / 2, cy: size / 2, r, fill: "none", stroke: "var(--ring-bg)", "stroke-width": 8 }, svg);
    const col = score >= 0.7 ? "var(--green)" : score >= 0.4 ? "var(--yellow)" : "var(--red)";
    el("circle", {
      cx: size / 2, cy: size / 2, r, fill: "none", stroke: col, "stroke-width": 8,
      "stroke-dasharray": `${(c * score).toFixed(1)} ${c.toFixed(1)}`,
      "stroke-linecap": "round", transform: `rotate(-90 ${size / 2} ${size / 2})`,
    }, svg);
    const t = el("text", { x: size / 2, y: size / 2 + 1, "text-anchor": "middle", class: "ring-val" }, svg);
    t.textContent = Math.round(score * 100) + "%";
    const l = el("text", { x: size / 2, y: size / 2 + 16, "text-anchor": "middle", class: "ring-lbl" }, svg);
    l.textContent = label || "";
  }

  return { bars, line, ring };
})();
