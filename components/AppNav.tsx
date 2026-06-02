"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Activity, LayoutDashboard, LineChart, Upload, Pill } from "lucide-react";

const LINKS = [
  { href: "/dashboard", label: "Command Center", icon: LayoutDashboard },
  { href: "/metrics", label: "Metrics", icon: LineChart },
  { href: "/data", label: "Data & Sources", icon: Upload },
  { href: "/records", label: "Meds & Conditions", icon: Pill },
];

export function AppNav({ name }: { name: string }) {
  const pathname = usePathname();
  return (
    <nav className="flex h-full flex-col">
      <div className="mb-6 flex items-center gap-2 px-2">
        <Activity className="h-6 w-6 text-emerald-600" />
        <span className="font-semibold">Command Center</span>
      </div>
      <ul className="space-y-1">
        {LINKS.map(({ href, label, icon: Icon }) => {
          const active = pathname === href || pathname.startsWith(href + "/");
          return (
            <li key={href}>
              <Link
                href={href}
                className={`flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition ${
                  active ? "bg-emerald-50 text-emerald-700" : "text-slate-600 hover:bg-slate-100"
                }`}
              >
                <Icon className="h-4 w-4" />
                {label}
              </Link>
            </li>
          );
        })}
      </ul>
      <div className="mt-auto px-2 pt-6 text-xs text-slate-400">{name}</div>
    </nav>
  );
}
