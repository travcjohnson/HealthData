"use client";

import {
  ResponsiveContainer, LineChart, Line, XAxis, YAxis, Tooltip,
  ReferenceArea, ReferenceLine, CartesianGrid,
} from "recharts";
import { format } from "date-fns";

export type ChartPoint = { t: number; v: number; label: string };

export function TrendChart({
  points,
  unit,
  refLow,
  refHigh,
  optimalLow,
  optimalHigh,
}: {
  points: ChartPoint[];
  unit: string;
  refLow: number | null;
  refHigh: number | null;
  optimalLow: number | null;
  optimalHigh: number | null;
}) {
  if (points.length === 0)
    return <div className="flex h-64 items-center justify-center text-slate-400">No data</div>;

  const ys = points.map((p) => p.v);
  const bounds = [refLow, refHigh, optimalLow, optimalHigh, ...ys].filter(
    (n): n is number => n != null,
  );
  const yMin = Math.min(...bounds);
  const yMax = Math.max(...bounds);
  const pad = (yMax - yMin) * 0.1 || 1;

  return (
    <ResponsiveContainer width="100%" height={320}>
      <LineChart data={points} margin={{ top: 10, right: 16, bottom: 0, left: -8 }}>
        <CartesianGrid strokeDasharray="3 3" stroke="#f1f5f9" />
        {/* optimal band (green) then reference band (amber edges) */}
        {optimalLow != null && optimalHigh != null && (
          <ReferenceArea y1={optimalLow} y2={optimalHigh} fill="#16a34a" fillOpacity={0.08} />
        )}
        {refLow != null && (
          <ReferenceLine y={refLow} stroke="#f59e0b" strokeDasharray="4 4" />
        )}
        {refHigh != null && (
          <ReferenceLine y={refHigh} stroke="#f59e0b" strokeDasharray="4 4" />
        )}
        <XAxis
          dataKey="t"
          type="number"
          domain={["dataMin", "dataMax"]}
          scale="time"
          tickFormatter={(t) => format(new Date(t), "MMM ''yy")}
          tick={{ fontSize: 11, fill: "#94a3b8" }}
          minTickGap={40}
        />
        <YAxis
          domain={[Math.floor(yMin - pad), Math.ceil(yMax + pad)]}
          tick={{ fontSize: 11, fill: "#94a3b8" }}
          width={44}
        />
        <Tooltip
          labelFormatter={(t) => format(new Date(Number(t)), "PP")}
          formatter={(v: number) => [`${v} ${unit}`, "value"]}
          contentStyle={{ fontSize: 12, borderRadius: 8, border: "1px solid #e2e8f0" }}
        />
        <Line
          type="monotone"
          dataKey="v"
          stroke="#0f766e"
          strokeWidth={2}
          dot={{ r: 2 }}
          activeDot={{ r: 4 }}
          isAnimationActive={false}
        />
      </LineChart>
    </ResponsiveContainer>
  );
}
