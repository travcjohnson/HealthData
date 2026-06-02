import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./lib/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        // semantic clinical status colors used across the dashboard
        status: {
          optimal: "#16a34a",
          normal: "#65a30d",
          borderline: "#d97706",
          high: "#dc2626",
          critical: "#991b1b",
          unknown: "#94a3b8",
        },
      },
      fontFamily: {
        sans: ["var(--font-sans)", "system-ui", "sans-serif"],
      },
    },
  },
  plugins: [],
};

export default config;
