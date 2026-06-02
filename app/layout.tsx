import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Health Command Center",
  description: "Your personal health dashboard — mixed data, clear trends.",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
